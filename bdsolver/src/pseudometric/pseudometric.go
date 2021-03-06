package pseudometric

import (
	"coupling"
	"markov"
	"sets"
	"matching"
	"setpair"
	"uvmethod"
	"disc"
	"utils"
	"log"
)

func InitD(n int) [][]float64{
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
		for j := range d[i] {
			if i != j {
				// when i and j is the same, we use the default value 0
				d[i][j] = 1
			}
		}
	}
	return d
}

func extractrandomfromset(tocompute [][]bool) (int, int) {
	for i := range tocompute {
		for j := range tocompute {
			if tocompute[i][j] == true {
				return i, j
			}
		}
	}
	panic("Tried to extract random element from empty set!")
}

func removeExactEdges(n *coupling.Node, exact [][]bool) {
	n.Visited = true
	
	for _, row := range n.Adj {
		for _, edge := range row {
			// if the edge node adj matrix is nill, we do not have to recursivevly call it,
			// and just just delete n from its successor slice
			if edge.To.Adj == nil {
				coupling.DeleteNodeInSlice(n, &edge.To.Succ)
				continue
			}
			
			// if the edge has already been visited, we do not have have to recursively call
			if edge.To.Visited {
				coupling.DeleteNodeInSlice(n, &edge.To.Succ)
				continue
			}
			
			// recursively removes edges and successor node bottom up
			removeExactEdges(edge.To, exact)
			coupling.DeleteNodeInSlice(n, &edge.To.Succ)
		}
	}
	
	// here we do the actual removing of edges
	exact[n.S][n.T] = true
	exact[n.T][n.S] = true
	n.Adj = nil
	n.Visited = false
	
	return
}

func PseudoMetric(m markov.MarkovChain, lambda float64, TPSolver func(markov.MarkovChain, *coupling.Node, [][]float64, float64, int, int)) [][]float64 {
	// initialize all the sets and the coupling
	n := len(m.Transitions)
	tocompute := sets.InitToCompute(len(m.Transitions))
	visited := sets.MakeMatrix(n)
	exact := sets.MakeMatrix(n)
	c := coupling.New()
	d := InitD(n)
	
	for !sets.EmptySet(tocompute) {
		var node *coupling.Node
		s, t := extractrandomfromset(tocompute)
		s, t = utils.GetMinMax(s, t)
		tocompute[s][t], tocompute[t][s] = false, false
		log.Printf("Run with node: (%v,%v)", s, t)
		
		if m.Labels[s] != m.Labels[t] {
			// s and t have the same label, so we set its distance to 0, and its exact and visited to true
			log.Printf("State %v and %v had different labels", s, t)
			d[s][t]= 1
			exact[s][t]= true
			visited[s][t] = true
			continue
		} else if s == t {
			// s and t are the same state, so we set its distance to 1, and its exact and visited to true
			log.Printf("State %v %v) was the same state", s, t)
			d[s][t] = 0
			exact[s][t] = true
			visited[s][t] = true
			continue
		} 
		
		// since the pair of states is not the same and share a label, we have to further process it
		// try to find the correpsonding node in the coupling
		node = coupling.FindNode(s, t, &c)
		
		if node == nil {
			// if FindNode returned a nil pointer, the node does not exist, and we have to make it
			node = matching.FindFeasibleMatching(m, s, t, &c)
			setpair.Setpair(m, node, exact, visited, d, &c)
		} else if node.Adj == nil {
			// if FindNode did return a node reference, but its adjacency matrix(matching) is nil, we have to create it
			node = matching.FindFeasibleMatching(m, s, t, &c)
			setpair.Setpair(m, node, exact, visited, d, &c)
		}
	
		disc.Disc(lambda, node, exact, d, &c)
		
		updateUntilOptimalSolutionsFound(lambda, m, node, exact, visited, d, c, TPSolver, []*coupling.Node{})
		
		removeExactEdges(node, exact)
		
		// remove everything that has been computed to exact, such that we will not try to solve it again
		tocompute = *sets.DifferensReal(&tocompute, &exact)
	}
	
	return d
}

func updateUntilOptimalSolutionsFound(lambda float64, m markov.MarkovChain, node *coupling.Node, exact [][]bool, visited [][]bool, d [][]float64, c coupling.Coupling, TPSolver func(markov.MarkovChain, *coupling.Node, [][]float64, float64, int, int), solvedNodes []*coupling.Node) {
	log.Printf("find optimal for: (%v,%v)", node.S, node.T)
	min, i, j := uvmethod.Run(node, d)
	// if min is negative, we can further improve it, so we update it using the TPSolver and iterated until we cannot improve it further
	for min < 0 && !utils.ApproxEqual(min, 0) {
		previ, prevj := i, j
		TPSolver(m, node, d, min, i, j)
		setpair.Setpair(m, node, exact, visited, d, &c)
		disc.Disc(lambda, node, exact, d, &c)
		
		min, i, j = uvmethod.Run(node, d)
		
		if previ == i && prevj == j && min < 0 {
			break
		}
	}
	
	// append solved nodes such that we do not end up recurively calling nodes that have already been found to be optimal
	solvedNodes = append(solvedNodes, node)
	
	for _, row := range node.Adj {
		for _, edge := range row {
			if edge.To.Adj == nil {
				// if the node do not have an adjacency matrix or is exact, we do not have to proccess it
				continue
			}
			if coupling.IsNodeInSlice(edge.To, solvedNodes) {
				// if the node has already been proccesses, we do not have to do it again
				continue
			}
			
			updateUntilOptimalSolutionsFound(lambda, m, edge.To, exact, visited, d, c, TPSolver, solvedNodes)
		}
	}
	
	exact[node.S][node.T] = true
	exact[node.S][node.T] = true
	
	return
}
