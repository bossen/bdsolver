package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sets"
)

func TestCorrectEdgesRemoved(t *testing.T) {
	c := setUpCouplingMatching()
	m := setUpMarkov()
	len := len(m.Transitions)
	visited := *sets.MakeMatrix(len)
	exact := *sets.MakeMatrix(len)
	d := make([][]float64, len, len)
	for i := 0; i < len; i++ {
		d[i] = make([]float64, len, len)
	}
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	
	removeExactEdges(n)
	
	assert.True(t, n.Adj == nil, "the adjacency matrix for (2,3) was not set to nil")
	assert.True(t, w.Adj != nil, "the adjacency matrix for (2,3) was set to nil")
	assert.False(t, w.Visited, "the visited bool for node (0,3) was not changed back to false")
	assert.False(t, n.Visited, "the visited bool for node (2,3) was not changed back to false")
}

func TestCorrectSuccNodesRemoved(t *testing.T) {
	c := setUpCouplingMatching()
	m := setUpMarkov()
	len := len(m.Transitions)
	visited := *sets.MakeMatrix(len)
	exact := *sets.MakeMatrix(len)
	d := make([][]float64, len, len)
	for i := 0; i < len; i++ {
		d[i] = make([]float64, len, len)
	}
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	n2 := n.Adj[1][2].To
	
	removeExactEdges(n)
	
	assert.True(t, succNode(w, n.Succ), "node (0,3) was removed as a successor for (2,3)")
	assert.False(t, succNode(n, w.Adj[0][0].To.Succ), "node (2,3) is still a successor for (0,1)")
	assert.True(t, succNode(w, w.Adj[0][0].To.Succ), "node (0,3) was removed as a successor for (0,1)")
	assert.False(t, succNode(n, n2.Succ), "node (2,3) was not removed as a successor for (2,2)")
}
