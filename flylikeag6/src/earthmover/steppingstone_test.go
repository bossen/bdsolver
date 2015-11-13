package earthmover

import (
	"coupling"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setUpCoupling() coupling.Coupling {
	c := coupling.New()
	n1 := coupling.Node{S: 0, T: 0}
	n2 := coupling.Node{S: 0, T: 1}
	n3 := coupling.Node{S: 1, T: 0}
	n4 := coupling.Node{S: 1, T: 1}
	e1 := coupling.Edge{&n1, 0.5, true}
	e2 := coupling.Edge{&n2, 0.2, true}
	e3 := coupling.Edge{&n3, 0, false}
	e4 := coupling.Edge{&n4, 0.3, true}
	n1.Succ = []*coupling.Node{&n1, &n2, &n4}
	n2.Succ = []*coupling.Node{&n2, &n4, &n1}
	n4.Succ = []*coupling.Node{&n2, &n1, &n4}
	n2.Adj = [][]*coupling.Edge{[]*coupling.Edge{&e1, &e2}, []*coupling.Edge{&e3, &e4}}
	c.Nodes = []*coupling.Node{&n1, &n2, &n3, &n4}

	return c
}

func TestSteppingStoneUpdateBasic(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	done := SteppingStone(n2, 1, 0)

	assert.True(t, done, "the stepping stone path were not completed")
	assert.True(t, n2.Adj[0][0].Basic, "cell did not remain a basic cell")
	assert.True(t, n2.Adj[0][1].Basic, "cell did not remain a basic cell")
	assert.True(t, n2.Adj[1][0].Basic, "cell did not become a basic cell")
	assert.False(t, n2.Adj[1][1].Basic, "cell did not become a non-basic cell")
}

func TestSteppingStoneReturnsFalse(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	done := SteppingStone(n2, 1, 1)

	assert.False(t, done, "a stepping stone path were found even though it does not exist")
	assert.True(t, n2.Adj[0][0].Basic, "cell were somehow changed to a non-basic cell")
	assert.True(t, n2.Adj[0][1].Basic, "cell were somehow changed to a non-basic cell")
	assert.False(t, n2.Adj[1][0].Basic, "cell were somehow changed to a basic cell")
	assert.True(t, n2.Adj[1][1].Basic, "cell were somehow changed to a non-basic cell")

	n2.Adj[1][1].Basic = false

	done = SteppingStone(n2, 1, 0)

	assert.False(t, done, "a stepping stone path were found even though it does not exist")
}

func TestpSteppingStoneRestoresVisited(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	SteppingStone(n2, 1, 0)

	assert.False(t, n2.Adj[0][0].To.Visited, "visited were not changed back to false")
	assert.False(t, n2.Adj[0][1].To.Visited, "visited were not changed back to false")
	assert.False(t, n2.Adj[1][0].To.Visited, "visited were not changed back to false")
	assert.False(t, n2.Adj[1][1].To.Visited, "visited were not changed back to false")
}

func TestGoHorizontalReturnsTrue(t *testing.T) {
	c := setUpCoupling()
	min := 2.0

	n := c.Nodes[1]

	done := goHorizontal(n, 1, 0, 1, 0, 2, 2, 1, &min)

	assert.True(t, done, "goHorizontal did not complete the stepping stone path")
}

func TestGoHorizontalReturnsFalse(t *testing.T) {
	c := setUpCoupling()
	min := 2.0

	n := c.Nodes[1]

	done := goHorizontal(n, 1, 1, 1, 1, 2, 2, 1, &min)

	assert.False(t, done, "goHorizontal somehow completed the stepping stone path")
}

func TestGoVerticalReturnsTrue(t *testing.T) {
	c := setUpCoupling()
	min := 0.3

	n := c.Nodes[1]

	done := goVertical(n, 1, 0, 1, 1, 2, 2, 2, &min)

	assert.True(t, done, "goVertical did not compplete the stepping stone path")
	assert.True(t, approxFloatEqual(0.0, n.Adj[1][0].Prob), "the cell (1 0) were changed")
	assert.False(t, approxFloatEqual(0.3, n.Adj[1][1].Prob), "the cell (1 1) were not changed")
}

func TestGoVerticalReturnsFalse(t *testing.T) {
	c := setUpCoupling()
	min := 2.0

	n := c.Nodes[1]

	done1 := goVertical(n, 1, 0, 1, 0, 2, 2, 1, &min)
	done2 := goVertical(n, 1, 1, 1, 1, 2, 2, 1, &min)

	assert.False(t, done1, "goVertical somehow completed the stepping stone path")
	assert.False(t, done2, "goVertical somehow completed the stepping stone path")
}

func TestCorrectSetSuccNodes(t *testing.T) {
	c := setUpCoupling()
	
	n2 := c.Nodes[1]
	
	SteppingStone(n2, 1, 0)
	
	assert.True(t, succNode(n2, c.Nodes[0].Succ), "node (1,0) did not remain a successor for (0,0)")
	assert.True(t, succNode(n2, c.Nodes[1].Succ), "node (1,0) did not remain a successor for (0,1)")
	assert.True(t, succNode(n2, c.Nodes[2].Succ), "node (1,0) did not become a successor for (1,0)")
	assert.False(t, succNode(n2, c.Nodes[3].Succ), "node (1,0) remained a successor for (1,1)")
}

func TestCorrectDeletedSuccNodes(t *testing.T) {
	c := setUpCoupling()
	
	succ := c.Nodes[0].Succ
	deleteSucc(succ[0], &succ)
	assert.Equal(t, len(succ), 2, "length of the successor slice did not change")
	deleteSucc(succ[0], &succ)
	assert.Equal(t, len(succ), 1, "length of the successor slice did not change")
}
