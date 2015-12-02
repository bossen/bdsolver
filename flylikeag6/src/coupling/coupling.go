package coupling

import (
	"log"
)

type Node struct {
	S, T       int
	Visited    bool
	Adj        [][]*Edge
	Succ       []*Node
	BasicCount int
}

type Edge struct {
	To    *Node
	Prob  float64
	Basic bool
}

type Coupling struct {
	Nodes []*Node
}

func New() Coupling {
	c := Coupling{}
	c.Nodes = make([]*Node, 0)
	return c
}

func FindNode(u, v int, c *Coupling) *Node {
	for _, n := range c.Nodes {
		if n.S == u && n.T == v {
			return n
		}
	}
	return nil
}

func Reachable(root *Node) []*Node {
	// Using slices might be slow. If we got performance problems we might
	// implement using lists instead.
	var reachables []*Node
	
	// Adding itself to reachables
	reachables = append(reachables, root)

	if root.Adj == nil {
		return reachables
	}
	
	root.Visited = true

	// Find all reachables from the  u,v node
	reachables = visit(root, reachables)

	log.Println("reachables:")
	for _, t := range reachables {
		log.Println(t)
	}

	for _, n := range reachables {
		n.Visited = false
	}

	return reachables
}

func visit(root *Node, results []*Node) []*Node {
	// log.Printf("%s, %s", root.S, root.T)
	if root.Adj == nil {
		return results
	}
	
	for i := range root.Adj {
		for j := range root.Adj[0] {
			edge := root.Adj[i][j]
			toVisit := edge.To
			
			if !edge.Basic || toVisit.Visited {
				continue
			} else if edge.Prob > 0 {
				toVisit.Visited = true
				results = append(results, toVisit)
				results = visit(toVisit, results)
			}	
		}
	}

	return results
}

func IsNodeInSlice(n *Node, nodes []*Node) bool {
	for _, succNode := range nodes {
		if succNode == n {
			return true
		}
	}
	return false
}

func DeleteNodeInSlice(n *Node, nodes *[]*Node) {
	for i := 0; i < len(*nodes); i++ {
		if (*nodes)[i] == n {
			// https://github.com/golang/go/wiki/SliceTricks
			(*nodes)[i] = (*nodes)[len(*nodes)-1]
			(*nodes)[len(*nodes)-1] = nil
			(*nodes) = (*nodes)[:len(*nodes)-1]
			break
		}
	}
	return
}
