package earthmover

import (
	"coupling"
	"log"
	"markov"
    "utils"
)

func findDemandedPairs(w *coupling.Node, visited [][]bool) []*coupling.Edge {
	demanded := make([]*coupling.Edge, 0)

	for _, rows := range w.Adj {
		for _, edge := range rows {
			if utils.ApproxEqual(edge.Prob, 0.0) {
				continue
			}
			if visited[edge.To.S][edge.To.T] || visited[edge.To.T][edge.To.S] {
				continue
			}

			demanded = append(demanded, edge)
		}
	}

	return demanded
}

func setpair(m markov.MarkovChain, w *coupling.Node, exact [][]bool, visited [][]bool, dist [][]float64, c *coupling.Coupling) {
	log.Printf("Setting pair for %v and %v", w.S, w.T)
	visited[w.S][w.T] = true

	for _, edge := range findDemandedPairs(w, visited) {
		u, v := utils.GetMinMax(edge.To.S, edge.To.T)

		visited[u][v] = true

		if u == v {
			log.Printf("%v and %v is the same state ", u, v)
			dist[u][v] = 0
			exact[u][v] = true
		} else if m.Labels[u] != m.Labels[v] {
			log.Printf("%v and %v have different labels ", u, v)
			dist[u][v]= 1
			exact[u][v]= true
		} else if !(exact[u][v] || exact[v][u]) {
			log.Printf("%v and %v have the same label ", u, v)
			w2 := findFeasibleMatching(m, u, v, c)
			setpair(m, w2, exact, visited, dist, c)
		}
	}
}
