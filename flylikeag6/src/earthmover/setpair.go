package earthmover

import (
	"coupling"
	"markov"
	"log"
)

func updatecoupling(c *coupling.Coupling, w [][]float64, s int, t int) int {
	return 1
}

func matchcardinality(w [][]float64) int {
	return 0
}

func nextdemandedpair(w [][]float64, i int) (int, int) {
	return 1, 1
}

func notexact(u int, v int, exact [][]bool) bool {
	return true
}

func setpair(m markov.MarkovChain, s int, t int, w [][]float64, exact [][]bool, visited [][]bool, dist [][]float64, c *coupling.Coupling) {
	updatecoupling(c, w, s, t) // this should be done in randommatching

	n := setUpNode() // all this should be changes once randomatching is done
	c.Nodes = []*coupling.Node{n.Adj[0][0].To, n.Adj[0][1].To, n.Adj[1][0].To, n.Adj[1][1].To}

	visited[s][t] = true
	visited[t][s] = true
	
	for _, rows := range n.Adj {
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
				setdistance(dist, u, v, 0)
				exact[v][u] = true
				
			} else if m.Labels[s] != m.Labels[t] {
				setdistance(dist, u, v, 1)
				exact[u][v] = true
				exact[v][u] = true
				
			} else if notexact(u, v, exact) {
				w2 := randommatching(m, u, v)
				setpair(m, u, v, w2, exact, visited, dist, c)
			}
		}
	}
}

func setUpNode() coupling.Node {
	n1 := coupling.Node{S: 0, T: 0}
	n2 := coupling.Node{S: 0, T: 1}
	n3 := coupling.Node{S: 1, T: 0}
	n4 := coupling.Node{S: 1, T: 1}
	e1 := coupling.Edge{&n1, 0.5, true}
	e2 := coupling.Edge{&n2, 0.2, true}
	e3 := coupling.Edge{&n3, 0, false}
	e4 := coupling.Edge{&n4, 0.3, true}
	n2.Adj = [][]*coupling.Edge{[]*coupling.Edge{&e1, &e2}, []*coupling.Edge{&e3, &e4}}
	
	return n2
}
