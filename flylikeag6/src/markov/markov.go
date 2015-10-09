package markov

func MakeMarkovChain() MarkovChain {
	n1 := Node{0, 1}
	n2 := Node{1, 1}
	nodes := []*Node{&n1, &n2}
	
	e1 := Edge{&n1, &n2, 0.3}
	e2 := Edge{&n2, &n1, 0.5}
	e3 := Edge{&n2, &n2, 0.5}
	e4 := Edge{&n1, &n1, 0.7}
	edges := []*Edge{&e1, &e2, &e3, &e4}
	
	var transitionMatrix [][]float64
	
	transitionMatrix = make([][]float64, 2, 2)
	for i := range transitionMatrix {
		transitionMatrix[i] = make([]float64, 2)
	}
	transitionMatrix[0][0] = 0.7 
	transitionMatrix[1][0] = 0.5
	transitionMatrix[0][1] = 0.3
	transitionMatrix[1][1] = 0.5
	
	return MarkovChain{nodes, edges, transitionMatrix}
}

type Node struct {
	Id, Label int
}

type Edge struct {
	Start, End *Node
	Prob float64
}

type MarkovChain struct {
	Nodes []*Node
	Edges []*Edge
	TransitionMatrix [][]float64
}
