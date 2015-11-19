package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
    "utils"
)

func TestCorrectRecursiveSetPairCall(t *testing.T) {
	// the same functions used for random matching testing
	c, m, visited, exact, d := setUpTest()

	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)

	assert.NotEqual(t, w.Adj, nil, "the adjacency matrix has not been filled for (0,3)")
	assert.NotEqual(t, w.Adj[2][2].To.Adj, nil, "the adjacency matrix has not been filled for (2,3)")
	assert.True(t, w.Adj[0][0].To.Adj == nil, "the adjacency matrix were somehow filled for (1,2)")
	assert.True(t, w.Adj[0][1].To.Adj == nil, "the adjacency matrix were somehow filled for (2,2)")
	assert.True(t, w.Adj[1][0].To.Adj == nil, "the adjacency matrix were somehow filled for (2,3)")
	// checks if the mutual node pointers in the two matchings are the same
	assert.Equal(t, w.Adj[0][0].To, w.Adj[2][2].To.Adj[0][0].To, "the nodes pointers were not the same for (1,2)")
	assert.Equal(t, w.Adj[0][1].To, w.Adj[2][2].To.Adj[0][1].To, "the nodes pointers were not the same for (2,2)")
	assert.Equal(t, w.Adj[0][2].To, w.Adj[2][2].To.Adj[0][2].To, "the nodes pointers were not the same for (2,3)")
	assert.Equal(t, w.Adj[1][0].To, w.Adj[2][2].To.Adj[1][0].To, "the nodes pointers were not the same for (1,3)")
	assert.Equal(t, w.Adj[1][1].To, w.Adj[2][2].To.Adj[1][1].To, "the nodes pointers were not the same for (2,3)")
	assert.Equal(t, w.Adj[1][2].To, w.Adj[2][2].To.Adj[1][2].To, "the nodes pointers were not the same for (3,3)")
}

func TestCorrectVisited(t *testing.T) {
	expected := [][]bool{
		[]bool{false, true, false, true, false, false, false},
		[]bool{true, true, true, false, false, false, false},
		[]bool{false, true, true, true, false, true, false},
		[]bool{true, false, true, false, false, false, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, true, false, false, false, false},
		[]bool{false, false, false, false, false, false, false}}

	c, m, visited, exact, d := setUpTest()

	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], visited[i][j], "the cell (%v,%v) were not correcly set", i, j)
		}
	}
}

func TestCorrectExact(t *testing.T) {
	expected := [][]bool{
		[]bool{false, true, false, false, false, false, false},
		[]bool{true, true, true, false, false, false, false},
		[]bool{false, true, true, false, false, true, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, true, false, false, false, false},
		[]bool{false, false, false, false, false, false, false}}

	c, m, visited, exact, d := setUpTest()

	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], exact[i][j], "the cell (%v,%v) were not correcly set", i, j)
		}
	}
}

func TestCorrectNestedMatchingFound(t *testing.T) {
	expected := [][]float64{
		[]float64{0.33, 0.17, 0.0},
		[]float64{0.0, 0.16, 0.34}}
	// the same functions used for random matching testing
	c, m, visited, exact, d := setUpTest()

	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)

	node := w.Adj[2][2].To

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.True(t, utils.ApproxEqual(expected[i][j], node.Adj[i][j].Prob), "the correct probability for cell (%v,%v) were not inserted", i, j)
		}
	}
}
func TestCorrectNestedBasicFound(t *testing.T) {
	expected := [][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true}}
	// the same functions used for random matching testing
	c, m, visited, exact, d := setUpTest()

	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)

	node := w.Adj[2][2].To

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], node.Adj[i][j].Basic, "the correct probability for cell (%v,%v) were not inserted", i, j)
		}
	}
}

func TestCorrectNestedSuccessorFound(t *testing.T) {
	// the same functions used for random matching testing
	c, m, visited, exact, d := setUpTest()

	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)

	node := w.Adj[2][2].To
	
	assert.True(t, succNode(node, node.Adj[0][0].To.Succ), "node (2,3) did not become a successor for (0,1)")
	assert.True(t, succNode(node, node.Adj[0][1].To.Succ), "node (2,3) did not become a successor for (1,1)")
	assert.True(t, succNode(node, node.Adj[1][1].To.Succ), "node (2,3) did not become a successor for (1,2)")
	assert.True(t, succNode(node, node.Adj[1][2].To.Succ), "node (2,3) did not become a successor for (2,2)")
	// common children between node(2,3) and w(0,3)
	assert.True(t, succNode(w, node.Adj[0][0].To.Succ), "node (0,1) did not have node (0,3) as a successor")
	assert.True(t, succNode(w, node.Adj[1][1].To.Succ), "node (1,2) did not have node (0,3) as a successor")
}
