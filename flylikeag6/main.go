package main

import (
	"coupling"
	"earthmover"
	"fmt"
	"log"
	"markov"
	"sets"
)

func TestUVmethod(node *coupling.Node, d *[256][256]float64) {
	log.Println(earthmover.Uvmethod(node, d))
}

//Computes all the possible combinations of the different nodes. This could be optimized, by setting everything below the i == j diagonal to false.
func initToCompute(n int) *[][]bool {
	toCompute := *sets.MakeMatrix(n)
	for i := range toCompute {
		for j := range toCompute {
			if i == j {
				toCompute[i][j] = false
			} else {
				toCompute[i][j] = true
			}
		}
	}
	return &toCompute
}
func main() {
	mymarkov := markov.New()
	tocompute := initToCompute(len(mymarkov.Transitions))
	earthmover.BipseudoMetric(mymarkov, 32, tocompute)
	// fmt.Printf("%+v", mymarkov)

	c := coupling.New()

	n1 := coupling.Node{S: 0, T: 0}
	n2 := coupling.Node{S: 0, T: 1}
	n3 := coupling.Node{S: 1, T: 0}
	n4 := coupling.Node{S: 1, T: 1}

	e1 := coupling.Edge{&n1, 0.5, true}
	e2 := coupling.Edge{&n2, 0.2, true}
	e3 := coupling.Edge{&n3, 0, false}
	e4 := coupling.Edge{&n4, 0.3, true}

	n2.Adj = &[][]coupling.Edge{[]coupling.Edge{e1, e2}, []coupling.Edge{e3, e4}}

	c.Nodes = []coupling.Node{n1, n2, n3, n4}

	fmt.Println(coupling.Reachable(0, 1, c))
  var d [256][256]float64
  d[0][0] = 5
  d[0][1] = 2
  d[1][0] = 0
  d[1][1] = 3
	TestUVmethod(&n2, &d)
}
