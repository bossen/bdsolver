package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"coupling"
)

func TestCorrectFilterNonZero(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	reachables := coupling.Reachable(w)
	nonZeroReachables := filterZeros(reachables, exact, d)
	
	assert.Equal(t, len(nonZeroReachables), 3, "the filtered node slice did not have length 3")
	assert.False(t, coupling.IsNodeInSlice(w.Adj[2][2].To, nonZeroReachables), "node (2,3) was not filtered")
	assert.True(t, coupling.IsNodeInSlice(w.Adj[0][0].To, nonZeroReachables), "node (0,1) was filtered")
	assert.True(t, coupling.IsNodeInSlice(w.Adj[3][2].To, nonZeroReachables), "node (2,5) was filtered")
	assert.False(t, coupling.IsNodeInSlice(w.Adj[2][2].To.Adj[0][1].To, nonZeroReachables), "node (1,1) was not filtered")
}

func TestCorrectReverseCouplingNodesFound(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	nonZeroReachables := []*coupling.Node{}
	
	nonZeroReachables = findReverseReachable(w.Adj[0][0].To, nonZeroReachables)
	
	assert.True(t, coupling.IsNodeInSlice(w, nonZeroReachables), "node (0,3) was not a successor for (0,1)")
	assert.True(t, coupling.IsNodeInSlice(w.Adj[2][2].To, nonZeroReachables), "node (2,3) was not a successor for (0,1)")
	assert.False(t, coupling.IsNodeInSlice(w.Adj[1][1].To, nonZeroReachables), "node (1,2) was a successor for (0,1)")
	assert.False(t, coupling.IsNodeInSlice(w.Adj[2][2].To.Adj[1][2].To, nonZeroReachables), "node (2,2) was a successor for (0,1)")
	
	nonZeroReachables2 := []*coupling.Node{}
	
	nonZeroReachables2 = findReverseReachable(w.Adj[2][2].To, nonZeroReachables2)
	
	assert.True(t, coupling.IsNodeInSlice(w, nonZeroReachables2), "node (0,3) was not a successor for (2,3)")
	assert.False(t, coupling.IsNodeInSlice(w.Adj[0][0].To, nonZeroReachables2), "node (0,1) was a successor for (2,3)")
}

func TestCorrectFindNonZeros(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	nonZero := findNonZero(w, exact, d, &c)
	
	assert.Equal(t, len(nonZero), 5, "the length of the non-zero node slice was not 5")
	assert.True(t, coupling.IsNodeInSlice(w, nonZero), "node (0,3) was not added to the non-zero slice")
	assert.True(t, coupling.IsNodeInSlice(w.Adj[2][2].To, nonZero), "node (2,3) was not added to the non-zero slice")
}
