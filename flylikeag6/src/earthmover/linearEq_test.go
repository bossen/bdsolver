package earthmover


import (
	"github.com/stretchr/testify/assert"
	"testing"
	"coupling"
	"utils"
)


func setUpLinearEquationFrame() ([][]float64, []float64, []NodePair) {
	a := make([][]float64, 1)
	a[0] = make([]float64, 1)
	a[0][0] = 1.0
	b := make([]float64, 1)
	index := make([]NodePair, 1)
	index[0] = NodePair{0, 3}
	
	return a, b, index
}

func TestFindsCorrectRowIndex(t *testing.T) {
	index := &[]NodePair{NodePair{0, 1}, NodePair{2, 4}, NodePair{0, 3}, NodePair{1, 7}}
	
	assert.Equal(t, findRowIndex(index, &coupling.Node{S: 0, T: 1}), 0, "the correct row index was not returned")
	assert.Equal(t, findRowIndex(index, &coupling.Node{S: 0, T: 3}), 2, "the correct row index was not returned")
	assert.Equal(t, findRowIndex(index, &coupling.Node{S: 1, T: 7}), 3, "the correct row index was not returned")
	assert.Equal(t, findRowIndex(index, &coupling.Node{S: 2, T: 4}), 1, "the correct row index was not returned")
}

func TestAddsNewLinearEquationCorrectly(t *testing.T) {
	a, b, index := setUpLinearEquationFrame()
	
	addLinearEquation(&a, &b, &index, &coupling.Node{S: 2, T: 3})
	
	// checking lengths of a, b, and index
	assert.Equal(t, len(a), 2, "the matrix did not a correct number of rows")
	assert.Equal(t, len(a[0]), 2, "the matrix did not a correct number of columns")
	assert.Equal(t, len(b), 2, "the vector did not a correct number of rows")
	assert.Equal(t, len(index), 2, "the slice did not a correct number of nodes")
	
	// checking if the diagonal of a are all 1.0
	assert.Equal(t, a[0][0], 1.0, "the value in matrix diagonal was not 1.0")
	assert.Equal(t, a[1][1], 1.0, "the value in matrix diagonal was not 1.0")
	
	addLinearEquation(&a, &b, &index, &coupling.Node{S: 1, T: 4})
	
	// checking lengths of a, b, and index
	assert.Equal(t, len(a), 3, "the matrix did not a correct number of rows")
	assert.Equal(t, len(a[0]), 3, "the matrix did not a correct number of columns")
	assert.Equal(t, len(b), 3, "the vector did not a correct number of rows")
	assert.Equal(t, len(index), 3, "the slice did not a correct number of nodes")
	
	// checking if the diagonal of a are all 1.0
	assert.Equal(t, a[0][0], 1.0, "the value in matrix diagonal was not 1.0")
	assert.Equal(t, a[1][1], 1.0, "the value in matrix diagonal was not 1.0")
	assert.Equal(t, a[2][2], 1.0, "the value in matrix diagonal was not 1.0")
}

func TestCorrectLinearFunctionsCreated(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	a, b, index := setUpLinearEquationFrame()
	
	setUpLinearEquations(w, exact, d, &a, &b, 0, &index, 1.0)
	
	// checking lengths of a, b, and index
	assert.Equal(t, len(a), 2, "the matrix did not a correct number of rows")
	assert.Equal(t, len(a[0]), 2, "the matrix did not a correct number of columns")
	assert.Equal(t, len(b), 2, "the vector did not a correct number of rows")
	assert.Equal(t, len(index), 2, "the slice did not a correct number of nodes")
	
	assert.True(t, utils.ApproxEqual(b[0], 0.83))
	assert.True(t, utils.ApproxEqual(b[1], 0.49))
	
	
	
	assert.Equal(t, a[0][0], 1.0, "the value in the matrix diagonal was not 1.0")
	assert.True(t, utils.ApproxEqual(a[0][1], -0.17), "the value in the matrix was not -0.17")
	assert.Equal(t, a[1][0], 0.0, "the value in the matrix was not 0.0")
	assert.Equal(t, a[1][1], 1.0, "the value in the matrix diagonal was not 1.0")
}

func TestCorrectLinearFunctionsCreatedLoops(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	a, b, index := setUpLinearEquationFrame()
	
	// inserting two loops into the coupling
	w.Adj[0][0] = &coupling.Edge{w, 0.33, true}
	w.Adj[2][2].To.Adj[1][2] = &coupling.Edge{w, 0.34, true}
	
	setUpLinearEquations(w, exact, d, &a, &b, 0, &index, 1.0)
	
	// checking lengths of a, b, and index
	assert.Equal(t, len(a), 2, "the matrix did not a correct number of rows")
	assert.Equal(t, len(a[0]), 2, "the matrix did not a correct number of columns")
	assert.Equal(t, len(b), 2, "the vector did not a correct number of rows")
	assert.Equal(t, len(index), 2, "the slice did not a correct number of nodes")
	
	assert.True(t, utils.ApproxEqual(b[0], 0.5))
	assert.True(t, utils.ApproxEqual(b[1], 0.49))
	
	assert.True(t, utils.ApproxEqual(a[0][0], 0.67), "the value in the matrix diagonal was not 0.67")
	assert.True(t, utils.ApproxEqual(a[0][1], -0.17), "the value in the matrix was not -0.17")
	assert.True(t, utils.ApproxEqual(a[1][0], -0.34), "the value in the matrix was not 0.0")
	assert.Equal(t, a[1][1], 1.0, "the value in the matrix diagonal was not 1.0")
}

