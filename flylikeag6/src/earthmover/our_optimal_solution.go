package earthmover

import (
	"coupling"
	"log"
)

func findOptimal(n *coupling.Node, d [][]float64) {
	m, i, j := Uvmethod(n, d)
	
	for m < 0 {
		res := SteppingStone(n, i, j)
		
		if !res {
			// stepping stone failed, try to recover basic nodes such that it can be completed
			recoverBasicNodes(n)
			
			res = SteppingStone(n, i, j)
			
			if !res {
				// still failed, something horrible is wrong
				log.Panic("stepping stone did not complete correctly")
			}
		}
		
		m, i, j = Uvmethod(n, d)
	}
}

type IntPair struct {
	I, J int
}

func recoverBasicNodes(n *coupling.Node) {
	var isolatedEdges []IntPair
	
	// checks every basic node if they are isolated in their row and column
	for i, row := range n.Adj {
		for j, edge := range row {
			if edge.Basic {
				checkIsolated(i, j, n, &isolatedEdges)
			}
		}
	}
	
	rowlen := len(n.Adj[0])
	
	for _, pair := range isolatedEdges {
		n.Adj[pair.I][(pair.J + 1) % rowlen].Basic = true
	}
}

func checkIsolated(s, t int, n *coupling.Node, isolatedEdges *[]IntPair) {
	for j := range(n.Adj[s]) {
		if j == t {
			continue
		}
		
		if n.Adj[s][j].Basic {
			// the basic node is not alone in its row
			return
		}
	}
	
	for i := range(n.Adj) {
		if i == s {
			continue
		}
		
		if n.Adj[i][t].Basic {
			// the basic node is not alone in its column
			return
		}
	}
	
	*isolatedEdges = append(*isolatedEdges, IntPair{s, t})
	
	return
}
