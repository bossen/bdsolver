package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsolatedEdgesFound(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	isolated := []IntPair{}
	
	checkIsolated(0, 0, w, &isolated)
	
	assert.Equal(t, len(isolated), 0, "somehow, node (0,1) were checked as isolated")
	
	checkIsolated(1, 1, w, &isolated)
	
	assert.Equal(t, len(isolated), 0, "somehow, node (1,2) were checked as isolated")
	
	checkIsolated(3, 2, w, &isolated)
	
	assert.Equal(t, len(isolated), 0, "somehow, node (2,5) were checked as isoalted")
	
	w.Adj[0][1].Basic = false
	w.Adj[1][2].Basic = false
	
	checkIsolated(0, 0, w, &isolated)
	
	assert.Equal(t, len(isolated), 1, "node (0,1) were not checked as isolated")
	
	checkIsolated(1, 1, w, &isolated)
	
	assert.Equal(t, len(isolated), 2, "node (1,2) were not checked as isolated")
}

func TestBasicNodesRecovered(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	w.Adj[0][1].Basic = false
	w.Adj[1][2].Basic = false
	
	res := SteppingStone(w, 2, 0)
	
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")

	recoverBasicNodes(w)
	
	assert.True(t, w.Adj[0][1].Basic, "node (1,1) were not set as basic")
	assert.True(t, w.Adj[1][2].Basic, "node (2,2) were not set as basic")

	res = SteppingStone(w, 2, 0)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 3, 0)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 2, 1)
	
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")
	
	recoverBasicNodes(w)
	
	res = SteppingStone(w, 2, 1)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
}
