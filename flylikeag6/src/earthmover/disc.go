package earthmover

import (
	"coupling"
	"log"
)

type NodePair struct {
	S, T int
}

func setZerosDistanceToZero(n *coupling.Node, nonzero []*coupling.Node, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	reachables := coupling.Reachable(n, c)
	
	for _, node := range nonzero {
		coupling.DeleteNodeInSlice(node, &reachables)
	}
	
	for _, node := range reachables {
		d[node.S][node.T] = 0
		exact[node.S][node.T] = true
	}
}

func setUpLinearEquations(n *coupling.Node, exact [][]bool, d [][]float64, a *[][]float64, b *[]float64, i int, index *[]NodePair) {
	n.Visited = true
	
	for _, row := range n.Adj {
		for _, edge := range row {
			// if the node is non-basic, we do not have to look at it
			if !edge.Basic {
				continue
			}
			
			// if the node is exact, we add its probablity times its distance to the i'th row in the b vector
			if exact[edge.To.S][edge.To.T] {
				(*b)[i] += d[edge.To.S][edge.To.T] * edge.Prob
				continue
			}
			
			// if the node has already been visited, we do not have to add it as a new linear equation
			// and subtract its probability from its corresponding place in the a matrix
			if edge.To.Visited {
				rowindex := findRowIndex(index, edge.To)
				(*a)[i][rowindex] -= edge.Prob
				continue
			}
			
			// we have to add a new linear equation, so we update a, b, and index
			addLinearEquation(a, b, index, edge.To)
			
			(*a)[i][len(*a)-1] -= edge.Prob
			
			// recursevely sets up the linear equation
			setUpLinearEquations(n, exact, d, a, b, len(*a)-1, index)
		}
	}
	
	return
}

func findRowIndex(index *[]NodePair, node *coupling.Node) int {
	for i, row := range *index {
		if node.S == row.S && node.T == row.T {
			return i
		}
	}
	panic("no row index was found")
}

func addLinearEquation(a *[][]float64, b *[]float64, index *[]NodePair, node *coupling.Node) {
	// b is a vector so we only need to append a new float 
	*b = append(*b, 0.0)
	
	// a is a matrix, so we have to append a new float to each inner slice ...
	for j := range(*a) {
		 (*a)[j] = append((*a)[j], 0.0)
	}
	
	rowlen := len((*a)[0])
	
	// ... and a new slice to the outer slice with the same length as the inner slices
	*a = append(*a, make([]float64, rowlen))
	
	(*a)[rowlen-1][rowlen-1] = 1.0
	
	// append a new NodePair to the index, since we have to manage a new linear equation
	*index = append(*index, NodePair{node.S, node.T})
}

func disc(lambda float64, n *coupling.Node, exact [][]bool, d[][]float64, c *coupling.Coupling) {
	log.Println("hello from disc.go!")
	nonzero := findNonZero(n, exact, d, c)
	
	setZerosDistanceToZero(n, nonzero, exact, d, c)
	
	// initial setup for the first linear equation, since there will always be at least one
	a := make([][]float64, 1)
	a[0] = make([]float64, 1)
	a[0][0] = 1.0
	b := make([]float64, 1)
	index := make([]NodePair, 1)
	index[0] = NodePair{n.S, n.T}
	
	// manipulates a, b, and index using pointers
	setUpLinearEquations(n, exact, d, &a, &b, 0, &index)
	
	// calculates the linear equations
	x, err := GaussPartial(a, b)
	
	if err != nil {
		log.Println("something horrible happened")
	}
	
	// uses index such that new distances are inserted in the correct places in the distance matrix
	for i, node := range index {
		d[node.S][node.T] = x[i]
	}
}
