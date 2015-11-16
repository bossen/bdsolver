package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sets"
	"coupling"
	"markov"
)

func setUpCouplingMatching() coupling.Coupling {
	return coupling.New()
}

func setUpMarkov() markov.MarkovChain {
	return markov.MarkovChain{
		Labels: []int{0, 1, 0, 0, 0, 1, 0},
		Transitions: [][]float64{
			[]float64{0.0, 0.33, 0.33, 0.17, 0.0, 0.17, 0.0},
			[]float64{0.0, 0.0, 0.4, 0.4, 0.0, 0.2, 0.0},
			[]float64{0.0, 0.5, 0.5, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.33, 0.33, 0.34, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.4, 0.4, 0.2, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.0, 0.1, 0.0, 0.2, 0.5, 0.2, 0.0},
			[]float64{0.0, 0.2, 0.33, 0.0, 0.1, 0.2, 0.17}}}
}

func setUpTest() (coupling.Coupling, markov.MarkovChain, [][]bool, [][]bool, [][]float64) {
	c := setUpCouplingMatching()
	m := setUpMarkov()
	n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := *sets.MakeMatrix(n)
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}
	
	return c, m, visited, exact, d
}

func TestCorrectEdgesRemoved(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	
	removeExactEdges(n, exact)
	
	assert.True(t, n.Adj == nil, "the adjacency matrix for (2,3) was not set to nil")
	assert.True(t, w.Adj != nil, "the adjacency matrix for (2,3) was set to nil")
	assert.False(t, w.Visited, "the visited bool for node (0,3) was not changed back to false")
	assert.False(t, n.Visited, "the visited bool for node (2,3) was not changed back to false")
}

func TestCorrectSuccNodesRemoved(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	n2 := n.Adj[1][2].To
	
	removeExactEdges(n, exact)
	
	assert.True(t, succNode(w, n.Succ), "node (0,3) was removed as a successor for (2,3)")
	assert.False(t, succNode(n, w.Adj[0][0].To.Succ), "node (2,3) is still a successor for (0,1)")
	assert.True(t, succNode(w, w.Adj[0][0].To.Succ), "node (0,3) was removed as a successor for (0,1)")
	assert.False(t, succNode(n, n2.Succ), "node (2,3) was not removed as a successor for (2,2)")
}

func TestCorrectExactSet(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	
	removeExactEdges(n, exact)
	
	assert.True(t, exact[n.S][n.T], "node (2,3) were not set to true in exact")
	assert.False(t, exact[w.S][w.T], "node (2,3) were not set to true in exact")
}

func TestCorrectFilterNonZero(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	r := coupling.Reachable(w.S, w.T, &c)
	r2 := filterZeros(r, exact, d)
	
	assert.Equal(t, len(r2), 3, "the filtered node slice did not have length 3")
	assert.False(t, succNode(w.Adj[2][2].To, r2), "node (2,3) was not filtered")
	assert.True(t, succNode(w.Adj[0][0].To, r2), "node (0,1) was filtered")
	assert.True(t, succNode(w.Adj[3][2].To, r2), "node (2,5) was filtered")
	assert.False(t, succNode(w.Adj[2][2].To.Adj[0][1].To, r2), "node (1,1) was not filtered")
}

func TestCorrectReverseCouplingNodesFound(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	r := []*coupling.Node{}
	
	r = findReverseReachable(w.Adj[0][0].To, r)
	
	assert.True(t, succNode(w, r), "node (0,3) was not a successor for (0,1)")
	assert.True(t, succNode(w.Adj[2][2].To, r), "node (2,3) was not a successor for (0,1)")
	assert.False(t, succNode(w.Adj[1][1].To, r), "node (1,2) was a successor for (0,1)")
	assert.False(t, succNode(w.Adj[2][2].To.Adj[1][2].To, r), "node (2,2) was a successor for (0,1)")
	
	r2 := []*coupling.Node{}
	
	r2 = findReverseReachable(w.Adj[2][2].To, r2)
	
	assert.True(t, succNode(w, r2), "node (0,3) was not a successor for (2,3)")
	assert.False(t, succNode(w.Adj[0][0].To, r2), "node (0,1) was a successor for (2,3)")
}

func TestCorrectFindNonZeros(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	r := findNonZero(w.S, w.T, exact, d, &c)
	
	assert.Equal(t, len(r), 5, "the length of the non-zero node slice was not 5")
	assert.True(t, succNode(w, r), "node (0,3) was not added to the non-zero slice")
	assert.True(t, succNode(w.Adj[2][2].To, r), "node (2,3) was not added to the non-zero slice")
}
