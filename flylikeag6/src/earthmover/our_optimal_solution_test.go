package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func TestOptimalSolutionFound(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	d = initD(len(m.Transitions))
	
	w := FindFeasibleMatching(m, 0, 3, &c)
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

func TestIsIntPairInSlice(t *testing.T) {
	intpairslice := []IntPair{IntPair{0, 1}, IntPair{4, 1}, IntPair{7, 3}, IntPair{2, 5}}
	
	assert.True(t, IsIntPairInSlice(IntPair{0, 1}, intpairslice), "the IntPair was not found")
	assert.True(t, IsIntPairInSlice(IntPair{4, 1}, intpairslice), "the IntPair was not found")
	assert.True(t, IsIntPairInSlice(IntPair{7, 3}, intpairslice), "the IntPair was not found")
	assert.True(t, IsIntPairInSlice(IntPair{2, 5}, intpairslice), "the IntPair was not found")
	
	assert.False(t, IsIntPairInSlice(IntPair{1, 5}, intpairslice), "the IntPair was found")
	assert.False(t, IsIntPairInSlice(IntPair{1, 0}, intpairslice), "the IntPair was found")
	assert.False(t, IsIntPairInSlice(IntPair{9, 9}, intpairslice), "the IntPair was found")
}

func TestTraverseVetical(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	w.Adj[1][2].Basic = false
	
	traversed := traverseVertical(w, []IntPair{}, 0)
	
	assert.Equal(t, 3, len(traversed), "the number of traversable basic nodes was not exactly 3")
	
	traversed = traverseVertical(w, []IntPair{}, 2)
	
	assert.Equal(t, 2, len(traversed), "the number of traversable basic nodes was not exactly 2")
	
	w.Adj[1][2].Basic = true
	
	traversed = traverseVertical(w, []IntPair{}, 2)
	
	assert.Equal(t, 6, len(traversed), "the number of traversable basic nodes was not exactly 6")
}

func TestTraverseHorizontal(t *testing.T) {
	c, m, _, _, _ := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	
	w.Adj[1][2].Basic = false
	
	traversed := traverseHorizontal(w, []IntPair{}, 0)
	
	assert.Equal(t, 3, len(traversed), "the number of traversable basic nodes was not exactly 3")
	
	traversed = traverseHorizontal(w, []IntPair{}, 2)
	
	assert.Equal(t, 2, len(traversed), "the number of traversable basic nodes was not exactly 2")
	
	w.Adj[1][2].Basic = true
	
	traversed = traverseHorizontal(w, []IntPair{}, 2)
	
	assert.Equal(t, 6, len(traversed), "the number of traversable basic nodes was not exactly 6")
}

func TestFindAllTraversableBasic(t *testing.T) {
	c, m, _, _, _ := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	
	w.Adj[1][2].Basic = false
	
	traversed := findAllTraversableBasic(w, IntPair{1, 2}, []IntPair{})
	
	assert.Equal(t, 5, len(traversed), "the number of traversable basic nodes was not exactly 5")
	
	w.Adj[1][2].Basic = true
	
	traversed = findAllTraversableBasic(w, IntPair{1, 2}, []IntPair{})
	
	assert.Equal(t, 6, len(traversed), "the number of traversable basic nodes was not exactly 6")
}

func TestFindFirstBasic(t *testing.T) {
	c, m, _, _, _ := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	
	index := findFirstBasic(w)
	
	assert.Equal(t, IntPair{0, 0}, index, "the index found for first basic was not correct")
	
	w.Adj[0][0].Basic = false
	w.Adj[0][1].Basic = false
	
	index = findFirstBasic(w)
	
	assert.Equal(t, IntPair{1, 1}, index, "the index found for first basic was not correct")
}

func TestBasicNodesRecovered(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	w.Adj[0][1].Basic = false
	w.Adj[1][2].Basic = false
	
	res := SteppingStone(w, 2, 0)
	
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")

	recoverBasicNodes(w)
	
	assert.True(t, w.Adj[0][1].Basic, "node (0,1) were not set as basic")
	assert.True(t, w.Adj[1][2].Basic, "node (1,2) were not set as basic")

	res = SteppingStone(w, 1, 0)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 3, 0)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 2, 1)
	
	assert.False(t, res, "stepping stone completed despite not enough basic nodes")
	
	recoverBasicNodes(w)
	
	assert.True(t, w.Adj[0][2].Basic, "node (0,1) were not set as basic")
	
	res = SteppingStone(w, 2, 1)
	
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
}

func TestBasicNodesRecovered2(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
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
	
	recoverBasicNodes(w)
	
	assert.True(t, w.Adj[0][2].Basic)
	
	res = SteppingStone(w, 1, 0)
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 2, 2)
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
	
	res = SteppingStone(w, 3, 1)
	assert.True(t, res, "stepping stone not completed despite enough basic nodes")
}
