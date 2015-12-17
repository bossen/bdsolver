package disc

import (
	"coupling"
	"log"
)

func setZerosDistanceToZero(n *coupling.Node, nonzero []*coupling.Node, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	reachables := coupling.Reachable(n)
	
	for _, node := range nonzero {
		coupling.DeleteNodeInSlice(node, &reachables)
	}
	
	for _, node := range reachables {
		d[node.S][node.T] = 0
		exact[node.S][node.T] = true
	}
}

func solveLinearEquations(n *coupling.Node, exact [][]bool, d [][]float64, lambda float64) ([]float64, []*coupling.Node) {
	// initial setup for the first linear equation, since there will always be at least one
	a := make([][]float64, 1)
	a[0] = make([]float64, 1)
	a[0][0] = 1.0
	b := make([]float64, 1)
	index := make([]*coupling.Node, 1)
	index[0] = n
	
	// manipulates a, b, and index using pointers
	setUpLinearEquations(n, exact, d, &a, &b, 0, &index, lambda)
	
	// solve the linear equations
	x, err := GaussPartial(a, b)
	
	if err != nil {
		log.Panic("it was not possible to calculate the linear updations for node (%v,%v)", n.S, n.T)
	}
	
	for _, node := range index {
		node.Visited = false
	}
	
	return x, index
}

func Disc(lambda float64, n *coupling.Node, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	log.Printf("tries to calculate linear equations for node (%v,%v)", n.S, n.T)
	//nonzero := findNonZero(n, exact, d, c)
	
	//setZerosDistanceToZero(n, nonzero, exact, d, c)
	
	x, index := solveLinearEquations(n, exact, d, lambda)
	
	// uses index such that new distances are inserted in the correct places in the distance matrix
	for i, node := range index {
		log.Printf("node (%v,%v)'s distance were set to: %v", node.S, node.T, x[i])
		d[node.S][node.T] = x[i]
	}
}
