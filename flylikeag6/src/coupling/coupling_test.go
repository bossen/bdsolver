package coupling

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func setUpCoupling() Coupling {
	c := New()
    n1 := Node{S: 0, T: 0}
    n2 := Node{S: 0, T: 1}
    n3 := Node{S: 1, T: 0}
    n4 := Node{S: 1, T: 1}
    e1 := Edge{&n1, 0.5, true}
    e2 := Edge{&n2, 0.2, true}
    e3 := Edge{&n3, 0, false}
    e4 := Edge{&n4, 0.3, true}
    n2.Adj = [][]*Edge{[]*Edge{&e1, &e2}, []*Edge{&e3, &e4}}
    c.Nodes = []*Node{&n1, &n2, &n3, &n4}
    
    return c
}

func TestStub(t *testing.T) {
	assert.True(t, true, "This is good. Canary test passing")
}

func TestReachableNilAdj(t *testing.T) {
	c := setUpCoupling()
    r := Reachable(0, 0, c)
    
    assert.Equal(t, len(r), 1, "size of the reachable set should be 1 if Adj is nil")
}

func TestResetsVisitedToFalse(t *testing.T) {
	c := setUpCoupling()
	
    // set the start to true just to check if it actually is set to false during reachable
    c.Nodes[1].Visited = true
    
    Reachable(0, 1, c)
    
    for i := 0; i < 4; i++ {
		assert.False(t, c.Nodes[i].Visited, "the visited variable date type was not changed back to false")
	}    
}

func TestReachableSetCardinality(t *testing.T) {
	c := setUpCoupling()
    
    r := Reachable(0, 1, c)
    
    assert.Equal(t, len(r), 3, "the cardinality of the reachable set should be 3")
    
    // test if it correctly finds the correct reachable set after we make one of the nodes unreachable
    c.Nodes[1].Adj[1][1].Prob = 0.0
    
    r = Reachable(0, 1, c)
    
    assert.Equal(t, len(r), 2, "the cardinality of the reachable set should be 2")
}
