package earthmover

import (
	"coupling"
	"log"
	"markov"
	"sets"
	"utils"
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

func findNode(s int, t int, c *coupling.Coupling) coupling.Node {
	newnode := coupling.Node{S: 0, T: 0}
	return newnode
}

func BipseudoMetric(m markov.MarkovChain, lambda int, tocompute *[][]bool) {
	// init
	n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := *sets.MakeMatrix(n)
	c := coupling.New()
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}

	for !sets.EmptySet(tocompute) {
		// Shouldn't we use utils.GetMinMax(s,t) here?
		s, t := extractrandomfromset(tocompute)
		s, t = utils.GetMinMax(s, t)
		log.Printf("Checking %v, %v", s, t)

		if m.Labels[s] != m.Labels[t] {
			d[s][t] = 1
			exact[s][t] = true
			visited[s][t] = true
		} else if s == t {
			d[s][t] = 0
			exact[s][t] = true
			visited[s][t] = true

		} else { // (s,t) is nontrivial.
			if !visited[s][t] {
				w := randomMatching(m, s, t, &c)
				setpair(m, s, t, w, exact, visited, d, &c)
			}

			disc(lambda, s, t, exact, d, &c)
			for _, node := range coupling.Reachable(s, t, &c) {

				// Optimal, so we dont need to check it.
				if node.Adj == nil {
					continue
				}

				for minimum, i, j := Uvmethod(node, d); minimum < 0; {
					// node := findNode(s, t, &c)
					SteppingStone(node, i, j)
				}

				// TODO this should be the optimal scheduling instead of random.
				w := randomMatching(m, s, t, &c)
				setpair(m, s, t, w, exact, visited, d, &c)
				disc(lambda, s, t, exact, d, &c)
			}

			// This could be done more prettier, e.g. have a reachable
			// function return a boolean matrix and use sets.UnionReal
			for _, r := range coupling.Reachable(s, t, &c) {
				exact[r.S][r.T] = true
			}
			// todo implement this:
			// removeedgesfromnodes(&c, &exact)
		}

		tocompute = sets.IntersectReal(*tocompute, exact)

		break //TODO remove this. This is for ending the code
	} // end for
	// setpair(m, 0, 1, w2, exact, visited, d, &c)
	// disc(1, 1, 1, exact, d, &c)
}
