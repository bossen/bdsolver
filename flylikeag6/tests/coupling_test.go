package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"coupling"
)

func TestStub(t *testing.T) {
	assert.True(t, true, "This is good. Canary test passing")
}

func TestReachableNilAdj(t *testing.T) {
	c := coupling.New()
    n := coupling.Node{S: 0, T: 0}
    c.Nodes = []*coupling.Node{&n}
    r := coupling.Reachable(0, 0, c)
    
    assert.Equal(t, len(r), 1, "size of the reachable set should be 1 if Adj is nil")
}

func TestResetsVisitedToFalse(t *testing.T) {
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
    
    // set the start to true just to check if it actually is set to false during reachable
    n2.Visited = true
    
    coupling.Reachable(0, 1, c)
    
    assert.False(t, n1.Visited, "the visited variable date type was not changed back to false")
    assert.False(t, n2.Visited, "the visited variable date type was not changed back to false")
    assert.False(t, n3.Visited, "the visited variable date type was not changed back to false")
    assert.False(t, n4.Visited, "the visited variable date type was not changed back to false")
}

func TestReachableSetCardinality(t *testing.T) {
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
    
    r := coupling.Reachable(0, 1, c)
    
    assert.Equal(t, len(r), 3, "the cardinality of the reachable set should be 3")
    
    // test if it correctly finds the correct reachable set after we make one of the nodes unreachable
    e4.Prob = 0.0
    r = coupling.Reachable(0, 1, c)
    
    assert.Equal(t, len(r), 2, "the cardinality of the reachable set should be 2")
}
