package earthmover

import (
	"coupling"
	"markov"
	"log"
)

func setpair(m markov.MarkovChain, s int, t int, w *coupling.Node, exact [][]bool, visited [][]bool, dist [][]float64, c *coupling.Coupling) {
	log.Printf("Setting pair for %v and %v", s, t)
	visited[s][t] = true
	visited[t][s] = true
	
	for _, rows := range w.Adj {
		for _, edge := range rows {
			if approxFloatEqual(edge.Prob, 0.0) {
				continue
			}
			
			u, v := edge.To.S, edge.To.T
			
			if u > v {
				temp := v
				v = u
				u = temp
			}
			
			if visited[u][v] || visited[v][u] {
				continue
			}
			
			visited[u][v] = true
			visited[v][u] = true
			
			if u == v {
				log.Printf("%v and %v is the same state ", u ,v)
				setdistance(dist, u, v, 0)
				exact[v][u] = true
				
			} else if m.Labels[u] != m.Labels[v] {
				log.Printf("%v and %v has different labels ", u, v)
				setdistance(dist, u, v, 1)
				exact[u][v] = true
				exact[v][u] = true
				
			} else if !(exact[u][v] || exact[v][u]) {
				log.Printf("%v and %v has the same label ", u, v)
				w2 := randomMatching(m, u, v, c)
				setpair(m, u, v, w2, exact, visited, dist, c)
			}
		}
	}
}
