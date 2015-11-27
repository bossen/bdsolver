package earthmover

import (
	"coupling"
	"log"
	"utils"
	"markov"
)

func FindOptimal(m markov.MarkovChain, n *coupling.Node, d [][]float64, min float64, i, j int) {
	for min < 0 {
		res := SteppingStone(n, i, j)
		
		if !res {
			log.Panic("halp!")
			log.Println("stepping stone failed, try to recover")
			recoverBasicNodes(n)
			
			res = SteppingStone(n, i, j)
			
			if !res {
				log.Panic("halp!")
			}			
			
			if !res {
				log.Println("stepping stone failed despite trying to recover")
				log.Panic("stepping stone did not complete correctly")
			}
		}
		
		recoverBasicNodes(n)
		
		min, i, j = Uvmethod(n, d)
	}
	
	log.Printf("node (%v,%v) is now optimal", n.S, n.T)
	return
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
		
		if n.Adj[s][j].Basic && !utils.ApproxEqual(n.Adj[s][j].Prob, 0) {
			// the basic node is not alone in its row
			return
		}
	}
	
	for i := range(n.Adj) {
		if i == s {
			continue
		}
		
		if n.Adj[i][t].Basic && !utils.ApproxEqual(n.Adj[i][t].Prob, 0) {
			// the basic node is not alone in its column
			return
		}
	}
	
	*isolatedEdges = append(*isolatedEdges, IntPair{s, t})
	
	return
}
