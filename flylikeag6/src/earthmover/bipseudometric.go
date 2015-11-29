package earthmover

import (
	"coupling"
	"markov"
	"sets"
	"utils"
	"log"
)

func initD(n int) [][]float64{
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

func BipseudoMetric(m markov.MarkovChain, lambda float64, TPSolver func(markov.MarkovChain, *coupling.Node, [][]float64, float64, int, int)) [][]float64 {
	// initialize all the sets and the coupling
	n := len(m.Transitions)
	tocompute := sets.InitToCompute(len(m.Transitions))
	visited := sets.MakeMatrix(n)
	exact := sets.MakeMatrix(n)
	c := coupling.New()
	d := initD(n)
	
	for !sets.EmptySet(tocompute) {
		s, t := extractrandomfromset(tocompute)
		s, t = utils.GetMinMax(s, t)
		tocompute[s][t], tocompute[t][s] = false, false
		log.Printf("Run with node: (%v,%v)", s, t)
		
		if m.Labels[s] != m.Labels[t] {
			// s and t have the same label, so we complete it and continue to the next one
			//log.Printf("State %v and %v had different labels", s, t)
			d[s][t], d[t][s] = 1, 1
			exact[s][t], exact[t][s] = true, true
			visited[s][t], visited[t][s] = true, true
			continue
		} else if s == t {
			// s and t are the same state, so we complete it and continue to the next one
			//log.Printf("State %v %v) was the same state", s, t)
			d[s][t] = 0
			exact[s][t] = true
			visited[s][t] = true
			continue
		}
		
		node := findFeasibleMatching(m, s, t, &c)
		setpair(m, node, exact, visited, d, &c)
		disc(lambda, node, exact, d, &c)
		
		findOptimalSolutions(lambda, m, node, exact, visited, d, c, TPSolver, []*coupling.Node{})
		
		removeExactEdges(node, exact)
		
		// remove everything that been computed to exact, such that we can not try to solve it again
		tocompute = *sets.DifferensReal(&tocompute, &exact)
	}
	for i := range d {
		log.Println(d[i])
	}
	return d
}

func findOptimalSolutions(lambda float64, m markov.MarkovChain, node *coupling.Node, exact [][]bool, visited [][]bool, d [][]float64, c coupling.Coupling, TPSolver func(markov.MarkovChain, *coupling.Node, [][]float64, float64, int, int), solvedNodes []*coupling.Node) {
	log.Printf("find optimal for: (%v,%v)", node.S, node.T)
	min, i, j := Uvmethod(node, d)
	
	// if min is negative, we cannot improve it further, so we skip using the TPSolver
	for min < 0 {
		TPSolver(m, node, d, min, i, j)
		setpair(m, node, exact, visited, d, &c)
		disc(lambda, node, exact, d, &c)
		
		min, i, j = Uvmethod(node, d)
	}
	
	// append solved nodes such that we to not end up recurively calling nodes that have already been found to be optimal
	solvedNodes = append(solvedNodes, node)
	
	for _, child := range coupling.Reachable(node) {
		if child.Adj == nil || exact[child.S][child.T] {
			// if the child do not have an adjacency matrix, he must already be exact
			continue
		}
		
		if coupling.IsNodeInSlice(child, solvedNodes) {
			// the child has already been solves, so we skip it
			continue
		}
		
		findOptimalSolutions(lambda, m, child, exact, visited, d, c, TPSolver, solvedNodes)
	}
	
	return
}
