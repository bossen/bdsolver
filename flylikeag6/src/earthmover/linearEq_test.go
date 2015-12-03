package earthmover


import (
	"github.com/stretchr/testify/assert"
	"testing"
	"coupling"
	"utils"
)

func setUpLinearEquationFrame() ([][]float64, []float64, []*coupling.Node) {
	a := make([][]float64, 1)
	a[0] = make([]float64, 1)
	a[0][0] = 1.0
	b := make([]float64, 1)
	index := make([]*coupling.Node, 1)
	index[0] = &coupling.Node{S: 0, T: 3}
	
	return a, b, index
}

func TestFindsCorrectRowIndex(t *testing.T) {
	index := &[]*coupling.Node{&coupling.Node{S: 0, T: 1}, &coupling.Node{S: 2, T: 4}, &coupling.Node{S: 0, T: 3}, &coupling.Node{S: 1, T: 7}}
	
	assert.Equal(t, 0, findRowIndex(index, &coupling.Node{S: 0, T: 1}), "the correct row index was not returned")
	assert.Equal(t, 2, findRowIndex(index, &coupling.Node{S: 0, T: 3}), "the correct row index was not returned")
	assert.Equal(t, 3, findRowIndex(index, &coupling.Node{S: 1, T: 7}), "the correct row index was not returned")
	assert.Equal(t, 1, findRowIndex(index, &coupling.Node{S: 2, T: 4}), "the correct row index was not returned")
}

func TestAddsNewLinearEquationCorrectly(t *testing.T) {
	a, b, index := setUpLinearEquationFrame()
	
	addLinearEquation(&a, &b, &index, &coupling.Node{S: 2, T: 3})
	
	// checking lengths of a, b, and index
	checkEquationDimensions(t, a, b, index, 2)
	
	// checking if the diagonal of a are all 1.0
	assert.Equal(t, 1.0, a[0][0], "the value in matrix diagonal was not 1.0")
	assert.Equal(t, 1.0, a[1][1], "the value in matrix diagonal was not 1.0")
	
	addLinearEquation(&a, &b, &index, &coupling.Node{S: 1, T: 4})
	
	checkEquationDimensions(t, a, b, index, 3)
	
	// checking if the diagonal of a are all 1.0
	assert.Equal(t, 1.0, a[0][0], "the value in matrix diagonal was not 1.0")
	assert.Equal(t, 1.0, a[1][1], "the value in matrix diagonal was not 1.0")
	assert.Equal(t, 1.0, a[2][2], "the value in matrix diagonal was not 1.0")
}

func TestCorrectLinearFunctionsCreated(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	a, b, index := setUpLinearEquationFrame()
	
	setUpLinearEquations(w, exact, d, &a, &b, 0, &index, 1.0)
	
	checkEquationDimensions(t, a, b, index, 2)
	
	assert.True(t, utils.ApproxEqual(b[0], 0.83))
	assert.True(t, utils.ApproxEqual(b[1], 0.49))
	
	assert.Equal(t, 1.0, a[0][0], "the value in the matrix diagonal was not 1.0")
	assert.True(t, utils.ApproxEqual(a[0][1], -0.17), "the value in the matrix was not -0.17")
	assert.Equal(t, 0.0, a[1][0], "the value in the matrix was not 0.0")
	assert.Equal(t, 1.0, a[1][1], "the value in the matrix diagonal was not 1.0")
}

func TestCorrectLinearFunctionsCreatedLoops(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	a, b, index := setUpLinearEquationFrame()
	
	// inserting two loops into the coupling
	w.Adj[0][0] = &coupling.Edge{w, 0.33, true}
	w.Adj[2][2].To.Adj[1][2] = &coupling.Edge{w, 0.34, true}
	
	setUpLinearEquations(w, exact, d, &a, &b, 0, &index, 1.0)
	
	checkEquationDimensions(t, a, b, index, 2)
	
	assert.True(t, utils.ApproxEqual(b[0], 0.5))
	assert.True(t, utils.ApproxEqual(b[1], 0.49))
	
	assert.True(t, utils.ApproxEqual(a[0][0], 0.67), "the value in the matrix diagonal was not 0.67")
	assert.True(t, utils.ApproxEqual(a[0][1], -0.17), "the value in the matrix was not -0.17")
	assert.True(t, utils.ApproxEqual(a[1][0], -0.34), "the value in the matrix was not 0.0")
	assert.Equal(t, 1.0, a[1][1], "the value in the matrix diagonal was not 1.0")
}

func checkEquationDimensions(t *testing.T, a [][]float64, b []float64, index []*coupling.Node, dim int) {
	// checking lengths of a, b, and index
	assert.Equal(t, dim, len(a), "the matrix did not have a correct number of rows")
	assert.Equal(t, dim, len(a[0]), "the matrix did not have a correct number of columns")
	assert.Equal(t, dim, len(b), "the vector did not have a correct number of rows")
	assert.Equal(t, dim, len(index), "the slice did not have a correct number of nodes")
}
