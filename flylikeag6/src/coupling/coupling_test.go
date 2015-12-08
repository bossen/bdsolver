package coupling

import (
	"testing"
	"github.com/stretchr/testify/assert"
//    "utils"
)

func TestReachableNilAdj(t *testing.T) {
	c := setUpCoupling()
    r := Reachable(c.Nodes[0])
    
    assert.Equal(t, len(r), 1, "size of the reachable set should be 1 if Adj is nil")
}

func TestResetsVisitedToFalse(t *testing.T) {
	c := setUpCoupling()
	
    // set the start to true just to check if it actually is set to false during reachable
    c.Nodes[1].Visited = true
    
    Reachable(c.Nodes[1])
    
    for i := 0; i < 4; i++ {
		assert.False(t, c.Nodes[i].Visited, "the visited variable date type was not changed back to false")
	}
	
	Reachable(c.Nodes[0])
	
	for i := 0; i < 4; i++ {
		assert.False(t, c.Nodes[i].Visited, "the visited variable date type was not changed back to false")
	}
}

func TestReachableSetCardinality(t *testing.T) {
	c := setUpCoupling()
    
    r := Reachable(c.Nodes[1])
    
    assert.Equal(t, len(r), 3, "the cardinality of the reachable set should be 3")
    
    // test if it correctly finds the correct reachable set after we make one of the nodes unreachable
    c.Nodes[1].Adj[1][1].Prob = 0.0
    
    r = Reachable(c.Nodes[1])
    
    assert.Equal(t, len(r), 2, "the cardinality of the reachable set should be 2")
}

func TestVisitRetunsIfAdjNil(t *testing.T) {
	c := setUpCoupling()
	n := c.Nodes[0]
	res := []*Node{n}
	
	r := visit(n, res)
	
	assert.Equal(t, len(r), 1, "visit function does not return immediately")
}

func TestVisitReturnsCorrectCardinality(t *testing.T) {
	c := setUpCoupling()
	n := c.Nodes[1]
	n.Visited = true
	res := []*Node{n}
	
	r := visit(n, res)
	
	assert.Equal(t, len(r), 3, "visit returns a result set that does not have size 3")
}


/*

func TestTraverseVetical(t *testing.T) {
	c, m, visited, exact, d := SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	setpair.Setpair(m, w, exact, visited, d, &c)
	
	w.Adj[1][2].Basic = false
	
	traversed := traverseVertical(w, []utils.IntPair{}, 0)
	
	assert.Equal(t, 3, len(traversed), "the number of traversable basic nodes was not exactly 3")
	
	traversed = traverseVertical(w, []utils.IntPair{}, 2)
	
	assert.Equal(t, 2, len(traversed), "the number of traversable basic nodes was not exactly 2")
	
	w.Adj[1][2].Basic = true
	
	traversed = traverseVertical(w, []utils.IntPair{}, 2)
	
	assert.Equal(t, 6, len(traversed), "the number of traversable basic nodes was not exactly 6")
}


func TestTraverseHorizontal(t *testing.T) {
	c, m, _, _, _ := coupling.SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	
	w.Adj[1][2].Basic = false
	
	traversed := traverseHorizontal(w, []utils.IntPair{}, 0)
	
	assert.Equal(t, 3, len(traversed), "the number of traversable basic nodes was not exactly 3")
	
	traversed = traverseHorizontal(w, []utils.IntPair{}, 2)
	
	assert.Equal(t, 2, len(traversed), "the number of traversable basic nodes was not exactly 2")
	
	w.Adj[1][2].Basic = true
	
	traversed = traverseHorizontal(w, []utils.IntPair{}, 2)
	
	assert.Equal(t, 6, len(traversed), "the number of traversable basic nodes was not exactly 6")
}

func TestFindAllTraversableBasic(t *testing.T) {
	c, m, _, _, _ := coupling.SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	
	w.Adj[1][2].Basic = false
	
	traversed := findAllTraversableBasic(w, utils.IntPair{1, 2}, []utils.IntPair{})
	
	assert.Equal(t, 5, len(traversed), "the number of traversable basic nodes was not exactly 5")
	
	w.Adj[1][2].Basic = true
	
	traversed = findAllTraversableBasic(w, utils.IntPair{1, 2}, []utils.IntPair{})
	
	assert.Equal(t, 6, len(traversed), "the number of traversable basic nodes was not exactly 6")
}

func TestFindFirstBasic(t *testing.T) {
	c, m, _, _, _ := coupling.SetUpTest()
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	
	index := findFirstBasic(w)
	
	assert.Equal(t, utils.IntPair{0, 0}, index, "the index found for first basic was not correct")
	
	w.Adj[0][0].Basic = false

}

*/
