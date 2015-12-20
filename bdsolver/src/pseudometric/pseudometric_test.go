package pseudometric

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"coupling"
	"matching"
	"setpair"
)

func TestCorrectEdgesRemoved(t *testing.T) {
	c, m, visited, exact, d := coupling.SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	setpair.Setpair(m, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	
	removeExactEdges(n, exact)
	
	assert.True(t, n.Adj == nil, "the adjacency matrix for (2,3) was not set to nil")
	assert.True(t, w.Adj != nil, "the adjacency matrix for (2,3) was set to nil")
	assert.False(t, w.Visited, "the visited bool for node (0,3) was not changed back to false")
	assert.False(t, n.Visited, "the visited bool for node (2,3) was not changed back to false")
}

func TestCorrectSuccNodesRemoved(t *testing.T) {
	c, m, visited, exact, d := coupling.SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	setpair.Setpair(m, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	n2 := n.Adj[1][2].To
	
	removeExactEdges(n, exact)
	
	assert.True(t, coupling.IsNodeInSlice(w, n.Succ), "node (0,3) was removed as a successor for (2,3)")
	assert.False(t, coupling.IsNodeInSlice(n, w.Adj[0][0].To.Succ), "node (2,3) is still a successor for (0,1)")
	assert.True(t, coupling.IsNodeInSlice(w, w.Adj[0][0].To.Succ), "node (0,3) was removed as a successor for (0,1)")
	assert.False(t, coupling.IsNodeInSlice(n, n2.Succ), "node (2,3) was not removed as a successor for (2,2)")
}

func TestCorrectExactSet(t *testing.T) {
	c, m, visited, exact, d := coupling.SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	setpair.Setpair(m, w, exact, visited, d, &c)
	
	n := w.Adj[2][2].To
	
	removeExactEdges(n, exact)
	
	assert.True(t, exact[n.S][n.T], "node (2,3) were not set to true in exact")
	assert.False(t, exact[w.S][w.T], "node (2,3) were not set to true in exact")
}

func TestInitializeD(t *testing.T) {
	d := InitD(100)
	for i := range d {
		for j := range d[i] {
			if i == j {
				assert.Zero(t, d[i][j], "d[%v][%v] value is %v", i, j, d[i][j])
			} else {
				assert.NotZero(t, d[i][j], "d[%v][%v] value is %v", i, j, d[i][j])
			}
		}
	}
}
