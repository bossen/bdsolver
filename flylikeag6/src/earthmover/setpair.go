package earthmover

import (
	"fmt"
    "sets"
    "markov"
)

func updatecoupling(coupling *int, w [][]float64, s int, t int) int {
	return 1
}

func matchcardinality(w [][]float64) int {
	return 0 
}

func nextdemandedpair(w [][]float64, i int) (int, int) {
	return 1, 1
}

func notexact(u int, v int, exact *int) bool {
	return true
}

func setpair(m markov.MarkovChain, s int, t int, w [][]float64, exact *int, visited *[][]bool, coupling *int) {
	fmt.Println("hi from setpair!")
	var _d [256][256] int
	d := &_d

	updatecoupling(coupling, w, s ,t)

	(*visited)[s][t] = true

	for i := 0; i < matchcardinality(w); i++ {
		u, v := nextdemandedpair(w, i)

	    (*visited)[u][v] = true

		if s == t {
			setdistance(d, u, v, 0)
			*exact = sets.UnionNode(*exact, u, v)
		} else if m.Labels[s] != m.Labels[t] {
			setdistance(d, u, v, 1)
			*exact = sets.UnionNode(*exact, u, v)
		} else if notexact(u, v, exact) {
			w2 := randommatching(m, u, v)
			setpair(m, u, v, w2, exact, visited, coupling)
		}
	}
}
