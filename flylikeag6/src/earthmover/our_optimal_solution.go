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
			recoverBasicNodes(n)
		}
		
		min, i, j = Uvmethod(n, d)
	}
	
	log.Printf("node (%v,%v) is now optimal", n.S, n.T)
	return
}

type IntPair struct {
	I, J int
}

func recoverBasicNodes(node *coupling.Node) {
	// the first basic node we have not traversed yet, such that we try and connect it with the other basic nodes
	firstbasic := findFirstBasic(node)
	
	// find all traversable to start the basic node recovery from
	t := findAllTraversableBasic(node, firstbasic, []IntPair{})
	
	recoverBasicNodesRecursive(node, t)
}

func recoverBasicNodesRecursive(node *coupling.Node, traversed []IntPair) {
	// create a copy such that we avoid out of bounds errors as we iterate
	t := traversed
	
	for _, pair := range t {
		// find index immediately left to the current index as i and j
		i, j := pair.I, (pair.J + 1) % len(node.Adj[0])
		
		if node.Adj[i][j].Basic {
			continue
		}
		
		// assuming index (i,j) is basic, we see what we can reach if it was
		tprime := findAllTraversableBasic(node, IntPair{i, j}, traversed)
		
		if len(tprime) == len(traversed) {
			// if we could not more more basic cells, we do not improve anything by setting (i,j) to basic
			continue
		}
		
		// otherwise, we set it to basic and make appropiate updates, especially updating traversed
		log.Printf("setting cell with index (%v,%v) in the matching for node (%v,%v)", i, j, node.S, node.T)
		node.Adj[i][j].Basic = true
		node.BasicCount++
		tprime = append(tprime, IntPair{i, j})
		traversed = tprime
		
		if len(traversed) == len(node.Adj) + (len(node.Adj[0]) - 1) {
			// if we have the correct number of nodes, we can now traverse all basic nodes, and we terminate
			return
		}
	}
	recoverBasicNodesRecursive(node, traversed)
	return
}

func findFirstBasic(node *coupling.Node) IntPair {
	for i, row := range node.Adj {
		for j, edge := range row {
			pair := IntPair{i, j}
			
			if !edge.Basic {
				continue
			}
			
			return pair
		}
	}
	return IntPair{-1, -1}
}

func findAllTraversableBasic(node *coupling.Node, curr IntPair, basicsfound []IntPair) []IntPair {
	if node.Adj[curr.I][curr.J].Basic {
		// if the current node is basic, add it to the result set
		// depending on where we call this function, we may or may not want to add the current node, this handles that
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
