package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"coupling"
)

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
