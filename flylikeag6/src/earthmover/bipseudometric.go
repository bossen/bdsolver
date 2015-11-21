package earthmover

import (
	"coupling"
	"log"
	"markov"
	"sets"
)

func extractrandomfromset(tocompute *[][]bool) (int, int) {
	for i := range *tocompute {
		for j := range *tocompute {
			if (*tocompute)[i][j] == true {
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
	n.Adj = nil
	n.Visited = false
	
	return
}

func getoptimalschedule(d [256][256]int, u int, v int) int {
	return 1
}

func isOptimal() bool {
	return true
}



func findNode(s int, t int, c *coupling.Coupling) coupling.Node {
	newnode := coupling.Node{S: 0, T: 0}
	return newnode
}

func BipseudoMetric(m markov.MarkovChain, lambda float64, tocompute *[][]bool) {
	n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := *sets.MakeMatrix(n)
	c := coupling.New()
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}

	w2 := findFeasibleMatching(m, 0, 1, &c)
	log.Println(w2)

	for !sets.EmptySet(tocompute) {
		s, t := extractrandomfromset(tocompute)
		log.Println(s)
		log.Println(t)

		if m.Labels[s] != m.Labels[t] {
			d[s][t] = 1
			exact[s][t] = true
			visited[s][t] = true
		} else if s == t {
			d[s][t] = 0
			exact[s][t] = true
			visited[s][t] = true

		} else {
			// if s,t not in visited ...

			node := findNode(s, t, &c)
			minimumvalue := 1.0 //TODO remove when
			_ = node
			//minimumvalue, iindex, jindex := Uvmethod(&node)
			for minimumvalue < 0 {
				//SteppingStone(&node, iindex, jindex)

				// w := getoptimalschedule(d, u, v) TODO this instead of next line
				w := findFeasibleMatching(m, s, t, &c)
				setpair(m, w, exact, visited, d, &c)
				disc(lambda, w, exact, d, &c)
				//minimumvalue, iindex, jindex = Uvmethod(&node)
			}
			//exact = sets.UnionReal(exact, reachable(s, t, c)) TODO update when reachable has been made
            // todo implement this:
            // removeedgesfromnodes(&c, &exact)
			removeExactEdges(w2, exact)
		}

		tocompute = sets.IntersectReal(*tocompute, *tocompute) //TODO THIS IS WRONG, use exact as second parameter, instead of tocompute twice

		break //TODO remove this. This is for ending the code
	}
	setpair(m, w2, exact, visited, d, &c)
	disc(1, w2, exact, d, &c)
	log.Println("hello world!")
	// return what?
}
