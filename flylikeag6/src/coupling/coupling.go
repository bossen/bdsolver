package coupling

import (
    "log"
)

type Node struct {
	S, T, Color int
	Adj [][]Edge
}

type Edge struct {
	To *Node
	Prob float64
	IsBasic bool
}

type Coupling struct {
	Nodes []Node
}

func New() Coupling {
    c := Coupling{}
    c.Nodes = make([]Node)
    return c
}


func Reachable(u, v int, c Coupling) []Node {
    // Using slices might be slow. If we got performance problems we might
    // implement using lists instead.
    var reachables []Node
    
    Root Node
    
    for i := range(c) {
		if c.Nodes[i].S == u && c.Nodes[i].T == v {
			Root = c.Nodes[i]
			break
		}
	}
	
	Root.Color = i

    // Adding itself to reachables
    reachables = append(reachables, Root)

    // Find all reachables from the  u,v node
    reachables = visit(u, v, c, reachables)

    log.Println("reachables:")
    for _, t := range reachables {
        log.Println(t)

    }

    for _, ce := range c.Matchings[StatePair{u, v}] {
        ce.Color = 0
    }
  return reachables
}


func visit(u, v int, c Coupling, results []Node)  []Node {
    for _, ce := range c.Matchings[StatePair{u, v}] {
        if ce.Color == 0 {
            ce.Color = 1
            results = append(results, StatePair{ce.S, ce.T})
            results = visit(ce.S, ce.T, c, results)
        }
    }
    return results
}
