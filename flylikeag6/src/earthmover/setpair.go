package earthmover

import (
	"coupling"
	"markov"
	"log"
)

func setpair(m markov.MarkovChain, s int, t int, w *coupling.Node, exact [][]bool, visited [][]bool, dist [][]float64, c *coupling.Coupling) {
	log.Println("s and t in setPair: ", s, t)
	log.Println("Visited: ", visited)
	visited[s][t] = true
	visited[t][s] = true
	log.Println("Visited: ", visited)
	
	for _, rows := range w.Adj {
		for _, edge := range rows {
			if approxFloatEqual(0, edge.Prob) {
				continue
			}
			
			u, v := edge.To.S, edge.To.T
			
			if visited[u][v] || visited[v][u] {
				continue
			}
			
			visited[u][v] = true
			visited[v][u] = true
			
			if u == v {
				log.Println("u and v is the same state ", u ,v)
				setdistance(dist, u, v, 0)
				exact[v][u] = true
				
			} else if m.Labels[s] != m.Labels[t] {
				log.Println("u and v has different labels ", u, v)
				setdistance(dist, u, v, 1)
				exact[u][v] = true
				exact[v][u] = true
				
			} else if !(exact[u][v] || exact[v][u]) {
				log.Println("u and v has the same label ", u, v)
				w2 := randomMatching(m, u, v, c)
				setpair(m, u, v, w2, exact, visited, dist, c)
			}
		}
	}
	log.Println("Visited: ", visited)
	log.Println("Exact: ", exact)
}
