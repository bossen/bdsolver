package main

import (
	"coupling"
	"earthmover"
	"log"
	"markov"
	"sets"
    "scanner"
)

func TestUVmethod(node *coupling.Node, d [][]float64) {
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


func testCompiler() {

    scanner, err := scanner.New("examples/algorithmfrompaper.lmc")
    if err != nil {
        log.Fatal("Not existing")
    }
    log.Println("Scanner reading word: ", scanner.ReadWord())
}

func main() {
    log.SetFlags(log.Lshortfile)
    
    

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

	n2.Adj = [][]*coupling.Edge{[]*coupling.Edge{&e1, &e2}, []*coupling.Edge{&e3, &e4}}

	c.Nodes = []*coupling.Node{&n1, &n2, &n3, &n4}

	d := make([][]float64, 2, 2)
	d[0] = make([]float64, 2, 2)
	d[1] = make([]float64, 2, 2)
	d[0][0] = 5.0
	d[0][1] = 2.0
	d[1][0] = 0.0
	d[1][1] = 3.0

	TestUVmethod(&n2, d)

	log.Println(coupling.Reachable(0, 1, &c))

	log.Println(n2.Adj[0][0], n2.Adj[0][1], n2.Adj[1][0], n2.Adj[1][1])
	earthmover.SteppingStone(&n2, 1, 0)
	log.Println(n2.Adj[0][0], n2.Adj[0][1], n2.Adj[1][0], n2.Adj[1][1])
    testCompiler()
}
