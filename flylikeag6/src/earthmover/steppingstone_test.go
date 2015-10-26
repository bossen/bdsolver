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
	n2.Adj = [][]*coupling.Edge{[]*coupling.Edge{&e1, &e2}, []*coupling.Edge{&e3, &e4}}
	c.Nodes = []*coupling.Node{&n1, &n2, &n3, &n4}

	return c
}

func TestSteppingStoneUpdateProbs(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	p1 := n2.Adj[0][0].Prob
	p2 := n2.Adj[0][1].Prob
	p3 := n2.Adj[1][0].Prob
	p4 := n2.Adj[1][1].Prob

	done := SteppingStone(n2, 1, 0)

	assert.True(t, done, "the stepping stone path were not completed")
	assert.False(t, approxFloatEqual(p1, n2.Adj[0][0].Prob), "coupling was not updated correctly")
	assert.False(t, approxFloatEqual(p2, n2.Adj[0][1].Prob), "coupling was not updated correctly")
	assert.False(t, approxFloatEqual(p3, n2.Adj[1][0].Prob), "coupling was not updated correctly")
	assert.False(t, approxFloatEqual(p4, n2.Adj[1][1].Prob), "coupling was not updated correctly")
}

func TestSteppingStoneUpdateBasic(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	done := SteppingStone(n2, 1, 0)

	b1 := n2.Adj[0][0].Basic
	b2 := n2.Adj[0][1].Basic
	b3 := n2.Adj[1][0].Basic
	b4 := n2.Adj[1][1].Basic

	assert.True(t, done, "the stepping stone path were not completed")
	assert.True(t, b1, "cell did not remain a basic cell")
	assert.True(t, b2, "cell did not remain a basic cell")
	assert.True(t, b3, "cell did not become a basic cell")
	assert.False(t, b4, "cell did not become a non-basic cell")
}

func TestSteppingStoneReturnsFalse(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	done := SteppingStone(n2, 1, 1)

	b1 := n2.Adj[0][0].Basic
	b2 := n2.Adj[0][1].Basic
	b3 := n2.Adj[1][0].Basic
	b4 := n2.Adj[1][1].Basic

	assert.False(t, done, "a stepping stone path were found even though it does not exist")
	assert.True(t, b1, "cell were somehow changed to a non-basic cell")
	assert.True(t, b2, "cell were somehow changed to a non-basic cell")
	assert.False(t, b3, "cell were somehow changed to a basic cell")
	assert.True(t, b4, "cell were somehow changed to a non-basic cell")

	n2.Adj[1][1].Basic = false

	done = SteppingStone(n2, 1, 0)

	assert.False(t, done, "a stepping stone path were found even though it does not exist")
}

func TestpSteppingStoneRestoresVisited(t *testing.T) {
	c := setUpCoupling()

	n2 := c.Nodes[1]

	SteppingStone(n2, 1, 0)

	v1 := n2.Adj[0][0].To.Visited
	v2 := n2.Adj[0][1].To.Visited
	v3 := n2.Adj[1][0].To.Visited
	v4 := n2.Adj[1][1].To.Visited

	assert.False(t, v1, "visited were not changed back to false")
	assert.False(t, v2, "visited were not changed back to false")
	assert.False(t, v3, "visited were not changed back to false")
	assert.False(t, v4, "visited were not changed back to false")
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
