package coupling

type Node struct {
	S, T    int
	Visited bool
	Adj     [][]*Edge
}

type Edge struct {
	To    *Node
	Prob  float64
	Basic bool
}

type Coupling struct {
	Nodes []*Node
}

func New() Coupling {
	c := Coupling{}
	c.Nodes = make([]*Node, 0)
	return c
}

func Reachable(u, v int, c Coupling) []*Node {
	// Using slices might be slow. If we got performance problems we might
	// implement using lists instead.
	var reachables []*Node

	var root *Node

	for _, n := range c.Nodes {
		if n.S == u && n.T == v {
			root = n
			break
		}
	}
	
	root.Visited = true
	
	// Adding itself to reachables
	reachables = append(reachables, root)

	if root.Adj == nil {
		return reachables
	}

	// Find all reachables from the  u,v node
	reachables = visit(root, reachables)

	for _, n := range c.Nodes {
		n.Visited = false
	}

	return reachables
}

func visit(root *Node, results []*Node) []*Node {
	// log.Printf("%s, %s", root.S, root.T)
	if root.Adj == nil {
		return results
	}
	for i := range root.Adj {
		for j := range root.Adj[0] {
			edge := root.Adj[i][j]
			toVisit := edge.To

			if edge.Prob > 0 && !(toVisit.Visited) {
				toVisit.Visited = true
				results = append(results, toVisit)
				results = visit(toVisit, results)
			}
		}
	}

	return results
}
