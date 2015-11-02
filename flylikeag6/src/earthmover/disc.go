package earthmover

import (
	"coupling"
	"fmt"
	"sets"
)

func reachable(u int, v int, graph coupling.Coupling) int {
	return 1
}

func transposegraph(graph coupling.Coupling) coupling.Coupling {
	return graph
}

func setdistance(d [][]float64, u int, v int, value float64) {
	d[u][v] = value
}

func putUnreachableInNonzero(s int, t int, c coupling.Coupling, nonzero *int, exact [][]bool, d [][]float64) {
	pairs := 1 //TODO remove when reachable has been implemented
	//pairs := sets.IntersectReal(reachable(s, t, coupling), exact) TODO update when reachable has been made
	pairsize := 1 //len(pairs) TODO
	for i := 0; i < pairsize; i++ {
		u, v := nextdemandedpairDisc(pairs, i)
		if d[u][v] > 0 {
			*nonzero = sets.Union(*nonzero, reachable(u, v, transposegraph(c)))
		}
	}
}

func setZerosDistanceToZero(s int, t int, nonzero int, exact [][]bool, d [][]float64, c coupling.Coupling) {
	pairs := sets.Differens(reachable(s, t, c), nonzero)
	pairsize := 1 //len(pairs) TODO
	for i := 0; i < pairsize; i++ {
		u, v := nextdemandedpairDisc(pairs, i)
		setdistance(d, u, v, 0)
		exact[u][v] = true
	}
}

func finda(c coupling.Coupling, nonzero int) int {
	return 1
}

func findb(exact [][]bool, d [][]float64, c coupling.Coupling, nonzero int) int {
	return 1
}

func solvelinearsystem(lambda int, a int, b int) int {
	return 1
}

func getvalue(x int, u int, v int) float64 {
	return 1.0
}

func updatedistances(nonzero int, d [][]float64, x int) {
	pairsize := 1 //len(x) TODO
	for i := 0; i < pairsize; i++ {
		u, v := nextdemandedpairDisc(nonzero, i)
		setdistance(d, u, v, getvalue(x, u, v))
	}
}

func nextdemandedpairDisc(w int, i int) (int, int) {
	return 1, 1
}

func disc(lambda int, s int, t int, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	nonzero := 1
	putUnreachableInNonzero(s, t, *c, &nonzero, exact, d)
	setZerosDistanceToZero(s, t, nonzero, exact, d, *c)
	a := finda(*c, nonzero)
	b := findb(exact, d, *c, nonzero)
	x := solvelinearsystem(lambda, a, b)
	updatedistances(nonzero, d, x)
	fmt.Println("discrepancy!")
}
