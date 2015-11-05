package earthmover

import (
    "sets"
    "testing"
	"github.com/stretchr/testify/assert"
)

func TestCorrectRecursiveSetPairCall(t *testing.T) {
	// the same functions used for random matching testing
	c := setUpCouplingMatching()
	m := setUpMarkov()
	n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := *sets.MakeMatrix(n)
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	assert.True(t, w.Adj != nil, "the adjacency matrix has not been filled for 0 3")
	assert.True(t, w.Adj[2][2].To.Adj != nil, "the adjacency matrix has not been filled for 2 3")
	assert.True(t, w.Adj[1][1].To.Adj == nil, "the adjacency matrix were somehow filled for node(%v,%v)", w.Adj[0][0].To.S, w.Adj[0][0].To.T)
	// checks if the mutual node pointers in the two matchings are the same
	assert.Equal(t, w.Adj[0][0].To, w.Adj[2][2].To.Adj[0][0].To, "the nodes pointers for (1,2) were not the same")
	assert.Equal(t, w.Adj[0][1].To, w.Adj[2][2].To.Adj[0][1].To, "the nodes pointers for (2,2) were not the same")
	assert.Equal(t, w.Adj[0][2].To, w.Adj[2][2].To.Adj[0][2].To, "the nodes pointers for (2,3) were not the same")
	assert.Equal(t, w.Adj[1][0].To, w.Adj[2][2].To.Adj[1][0].To, "the nodes pointers for (1,3) were not the same")
	assert.Equal(t, w.Adj[1][1].To, w.Adj[2][2].To.Adj[1][1].To, "the nodes pointers for (2,3) were not the same")
	assert.Equal(t, w.Adj[1][2].To, w.Adj[2][2].To.Adj[1][2].To, "the nodes pointers for (3,3) were not the same")
}

func TestCorrectNestedMatchingFound(t *testing.T) {
	expected := [][]float64{
		[]float64{0.33, 0.17, 0.0},
		[]float64{0.0, 0.16, 0.34}}
	// the same functions used for random matching testing
	c := setUpCouplingMatching()
	m := setUpMarkov()
	n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := *sets.MakeMatrix(n)
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	node := w.Adj[2][2].To
	
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.True(t, approxFloatEqual(expected[i][j], node.Adj[i][j].Prob), "the correct probability were not inserted")
		}
	}
}
func TestCorrectNestedBasicFound(t *testing.T) {
	expected := [][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true}}
	// the same functions used for random matching testing
	c := setUpCouplingMatching()
	m := setUpMarkov()
	n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := *sets.MakeMatrix(n)
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	node := w.Adj[2][2].To
	
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], node.Adj[i][j].Basic, "the correct probability were not inserted")
		}
	}
}

