package earthmover

import (
	"coupling"
	"log"
)

func setdistance(d [][]float64, u int, v int, value float64) {
 	d[u][v] = value
 }

func setZerosDistanceToZero(n *coupling.Node, nonzero []*coupling.Node, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	r := coupling.Reachable(n, c)
	
	for _, node := range nonzero {
		coupling.DeleteNodeInSlice(node, &r)
	}
	
	for _, node := range r {
		d[node.S][node.T] = 0
		exact[node.S][node.T] = true
	}
}

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

func disc(lambda int, n *coupling.Node, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	nonzero := findNonZero(n, exact, d, c)
	
	setZerosDistanceToZero(n, nonzero, exact, d, c)
	
	a := finda(*c, nonzero)
	b := findb(exact, d, *c, nonzero)
	x := solvelinearsystem(lambda, a, b)
	updatedistances(nonzero, d, x)
	log.Println("discrepancy!")
}
