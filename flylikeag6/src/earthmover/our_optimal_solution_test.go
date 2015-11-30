package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func TestOptimalSolutionFound(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	d = initD(len(m.Transitions))
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	// the expected solution is not precice since the markov chain used do not use precise fractions
	expected := [][]float64{
		[]float64{0, 0.33, 0},
		[]float64{0, 0, 0.33},
		[]float64{0.17, 0, 0},
		[]float64{0.16, 0, 0.01}}
		
	min, i, j := Uvmethod(w, d)
	
	FindOptimal(m, w, d, min, i, j)
	
	for i := range(w.Adj) {
		for j := range(w.Adj[0]) {
			assert.True(t, utils.ApproxEqual(expected[i][j], w.Adj[i][j].Prob), "the optimal probability found was not what we expected")
		}
	}
}

func TestBasicNodesRecovered(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	w.Adj[0][1].Basic = false
	w.Adj[1][2].Basic = false
	
	res := SteppingStone(w, 2, 0)
	
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")

	recoverBasicNodes(w, []IntPair{})
	
	assert.True(t, w.Adj[0][1].Basic, "node (1,1) were not set as basic")
	assert.True(t, w.Adj[2][0].Basic, "node (2,2) were not set as basic")

	res = SteppingStone(w, 1, 0)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 3, 0)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 2, 1)
	
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")
	
	recoverBasicNodes(w, []IntPair{})
	
	res = SteppingStone(w, 2, 1)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
}

func TestBasicNodesRecovered2(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	w.Adj[1][1].Basic = false
	w.Adj[1][1].Prob = 0
	w.Adj[2][1].Basic = true
	w.Adj[2][1].Prob = 0.33
	w.Adj[1][2].Basic = true
	w.Adj[1][2].Prob = 0.17
	w.Adj[2][2].Basic = false
	w.Adj[2][2].Prob = 0
	w.BasicCount = 5
	
	res := SteppingStone(w, 1, 0)
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")
	
	res = SteppingStone(w, 2, 2)
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")
	
	res = SteppingStone(w, 3, 1)
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")
	
	recoverBasicNodes(w, []IntPair{})
	
	assert.True(t, w.Adj[0][2].Basic)
	
	res = SteppingStone(w, 1, 0)
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 2, 2)
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 3, 1)
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
}
