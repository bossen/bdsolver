package main

import (
	"coupling"
	"earthmover"
	"log"
    "scanner"
    "markov"
)

func TestUVmethod(node *coupling.Node, d [][]float64) {
	log.Println(earthmover.Uvmethod(node, d))
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

	mymarkov := markov.MarkovChain{
		Labels: []int{0, 1, 0, 0, 0, 1, 0},
		Transitions: [][]float64{
			[]float64{0.0, 0.33333, 0.33333, 0.16667, 0.0, 0.16667, 0.0},
			[]float64{0.0, 0.0, 0.4, 0.4, 0.0, 0.2, 0.0},
			[]float64{0.0, 0.5, 0.5, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.33333, 0.33333, 0.33334, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.4, 0.4, 0.2, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.0, 0.1, 0.0, 0.2, 0.5, 0.2, 0.0},
			[]float64{0.0, 0.2, 0.33333, 0.0, 0.1, 0.2, 0.16667}}}
	earthmover.BipseudoMetric(mymarkov, 0.95, earthmover.FindOptimal)
	
	
	
	
	// fmt.Printf("%+v", mymarkov)
	/*
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

	log.Println(coupling.Reachable(&n2))

	log.Println(n2.Adj[0][0], n2.Adj[0][1], n2.Adj[1][0], n2.Adj[1][1])
	earthmover.SteppingStone(&n2, 1, 0)
	log.Println(n2.Adj[0][0], n2.Adj[0][1], n2.Adj[1][0], n2.Adj[1][1])
    testCompiler()
	
	a := [][]float64{[]float64{1.0, -(1.0/6.0)}, []float64{0.0, 1.0}}
	b := []float64{5.0/6.0, 1.0/2.0}
	x, err := earthmover.GaussPartial(a, b)
	
	log.Println(a)
	log.Println(b)
	log.Println(x)
	log.Println(err)
	*/
}
