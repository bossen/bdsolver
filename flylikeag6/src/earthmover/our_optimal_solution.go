package earthmover

import (
	"coupling"
	"log"
	"markov"
)

func FindOptimal(m markov.MarkovChain, n *coupling.Node, d [][]float64, min float64, i, j int) {
	for min < 0 {
		SteppingStone(n, i, j)
		
		if n.BasicCount < len(n.Adj) + (len(n.Adj[0]) - 1) {
			log.Printf("Recover basic nodes for node (%v,%v)", n.S, n.T)
			recoverBasicNodes(n, []IntPair{})
		}
		
		min, i, j = Uvmethod(n, d)
	}
	
	log.Printf("node (%v,%v) is now optimal", n.S, n.T)
	return
}

type IntPair struct {
	I, J int
}

func recoverBasicNodes(node *coupling.Node, traversed []IntPair) {
	firstbasic := findFirstNonTraversedBasic(node, traversed)
	
	t := findAllTraversableBasic(node, firstbasic, []IntPair{})
	traversed = t
	
	for _, pair := range t {
		i, j := pair.I, (pair.J + 1) % len(node.Adj[0])
		
		if node.Adj[i][j].Basic {
			continue
		}
		
		tprime := findAllTraversableBasic(node, IntPair{i, j}, traversed)
		
		if len(tprime) == len(traversed) {
			continue
		}
		
		node.Adj[i][j].Basic = true
		node.BasicCount++
		tprime = append(tprime, IntPair{i, j})
		traversed = tprime
		
		log.Println(len(node.Adj) + (len(node.Adj[0]) - 1))
		log.Println(len(traversed))
		
		if len(traversed) == len(node.Adj) + (len(node.Adj[0]) - 1) {
			return
		}
	}
	recoverBasicNodes(node, traversed)
}

func findFirstNonTraversedBasic(node *coupling.Node, traversed []IntPair) IntPair {
	for i, row := range node.Adj {
		for j, edge := range row {
			pair := IntPair{i, j}
			
			if !edge.Basic || IsIntPairInSlice(pair, traversed){
				continue
			}
			
			return pair
		}
	}
	return IntPair{-1, -1}
}

func findAllTraversableBasic(node *coupling.Node, curr IntPair, basicsfound []IntPair) []IntPair {
	if node.Adj[curr.I][curr.J].Basic {
		basicsfound = append(basicsfound, curr)
	}
	
	basicsfound = traverseHorizontal(node, basicsfound, curr.I)
	basicsfound = traverseVertical(node, basicsfound, curr.J)
	
	return basicsfound
}

func traverseHorizontal(node *coupling.Node, basicsfound []IntPair, i int) []IntPair {
	for j := range node.Adj[0] {
		pair := IntPair{i, j}
		
		if !node.Adj[i][j].Basic || IsIntPairInSlice(pair, basicsfound) {
			continue
		}
		
		basicsfound = append(basicsfound, pair)
		
		basicsfound = traverseVertical(node, basicsfound, j)
	}
	
	return basicsfound
}

func traverseVertical(node *coupling.Node, basicsfound []IntPair, j int) []IntPair {
	for i := range node.Adj {
		pair := IntPair{i, j}
		
		if !node.Adj[i][j].Basic || IsIntPairInSlice(pair, basicsfound) {
			continue
		}
		
		basicsfound = append(basicsfound, pair)
		
		basicsfound = traverseHorizontal(node, basicsfound, i)
	}
	
	return basicsfound
}

func IsIntPairInSlice(pair IntPair, pairs []IntPair) bool {
	for _, n := range pairs {
		if n == pair {
			return true
		}
	}
	return false
}
