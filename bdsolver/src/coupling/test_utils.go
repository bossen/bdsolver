package coupling

import (
	"markov"
	"sets"
)



func SetUpMarkov() markov.MarkovChain {
	return markov.MarkovChain{
		Labels: []int{0, 1, 0, 0, 0, 1, 0},
		Transitions: [][]float64{
			[]float64{0.0, 0.33, 0.33, 0.17, 0.0, 0.17, 0.0},
			[]float64{0.0, 0.0, 0.4, 0.4, 0.0, 0.2, 0.0},
			[]float64{0.0, 0.5, 0.5, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.33, 0.33, 0.34, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.4, 0.4, 0.2, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.0, 0.1, 0.0, 0.2, 0.5, 0.2, 0.0},
			[]float64{0.0, 0.2, 0.33, 0.0, 0.1, 0.2, 0.17}}}
}

func SetUpTest() (Coupling, markov.MarkovChain, [][]bool, [][]bool, [][]float64) {
	c := New()
	m := SetUpMarkov()
	n := len(m.Transitions)
	visited := sets.MakeMatrix(n)
	exact := sets.MakeMatrix(n)
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
	}
	
	return c, m, visited, exact, d
}

func setUpCoupling() Coupling {
	c := New()
    n1 := Node{S: 0, T: 0}
    n2 := Node{S: 0, T: 1}
    n3 := Node{S: 1, T: 0}
    n4 := Node{S: 1, T: 1}
    e1 := Edge{&n1, 0.5, true}
    e2 := Edge{&n2, 0.2, true}
    e3 := Edge{&n3, 0, false}
    e4 := Edge{&n4, 0.3, true}
    n2.Adj = [][]*Edge{[]*Edge{&e1, &e2}, []*Edge{&e3, &e4}}
    c.Nodes = []*Node{&n1, &n2, &n3, &n4}
    
    return c
}
