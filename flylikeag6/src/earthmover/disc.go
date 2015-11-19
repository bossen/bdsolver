package earthmover

import (
	"coupling"
	"fmt"
)

func transposegraph(graph coupling.Coupling) coupling.Coupling {
	return graph
}

func setdistance(d [][]float64, u int, v int, value float64) {
	d[u][v] = value
}

/*
func setZerosDistanceToZero(s int, t int, nonzero int, exact [][]bool, d [][]float64, c coupling.Coupling) {
	pairs := sets.Differens(reachable(s, t, c), nonzero)
	pairsize := 1 //len(pairs) TODO
	for i := 0; i < pairsize; i++ {
		u, v := nextdemandedpairDisc(pairs, i)
		setdistance(d, u, v, 0)
		exact[u][v] = true
	}
}
* */

func finda(c coupling.Coupling, nonzero []*coupling.Node) int {
	return 1
}

func findb(exact [][]bool, d [][]float64, c coupling.Coupling, nonzero []*coupling.Node) int {
	return 1
}

func solvelinearsystem(lambda int, a int, b int) int {
	return 1
}

func getvalue(x int, u int, v int) float64 {
	return 1.0
}

func updatedistances(nonzero []*coupling.Node, d [][]float64, x int) {
	pairsize := 1 //len(x) TODO
	for i := 0; i < pairsize; i++ {
		u, v := nextdemandedpairDisc(nonzero, i)
		setdistance(d, u, v, getvalue(x, u, v))
	}
}

func nextdemandedpairDisc(w []*coupling.Node, i int) (int, int) {
	return 1, 1
}

func disc(lambda int, s int, t int, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	nonzero := findNonZero(s, t, exact, d, c)
	//setZerosDistanceToZero(s, t, nonzero, exact, d, *c)
	a := finda(*c, nonzero)
	b := findb(exact, d, *c, nonzero)
	x := solvelinearsystem(lambda, a, b)
	updatedistances(nonzero, d, x)
	fmt.Println("discrepancy!")
}
