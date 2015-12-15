package tpsolverdefault

import (
	"coupling"
	"log"
	"markov"
    "uvmethod"
    "utils"
)

func Solve(m markov.MarkovChain, n *coupling.Node, d [][]float64, min float64, i, j int) {
	log.Println("Running default optimizer")
	for min < 0 && utils.ApproxEqual(min, 0) {
		SteppingStone(n, i, j)
		
		if n.BasicCount < len(n.Adj) + (len(n.Adj[0]) - 1) {
			log.Printf("Recover basic nodes for node (%v,%v)", n.S, n.T)
			coupling.RecoverBasicNodes(n)
		}
		
		min, i, j = uvmethod.Run(n, d)
	}
	
	log.Printf("node (%v,%v) is now optimal", n.S, n.T)
	return
}
