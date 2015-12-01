package earthmover

import (
	"coupling"
	"math"
    "utils"
    "log"
)

func SteppingStone(n *coupling.Node, s int, t int) bool {
	log.Printf("try to find path for cell: (%v,%v) with index (%v,%v) in matching for node", s, t, n.S, n.T)
	min, rowBound, colBound := 2.0, len(n.Adj), len(n.Adj[0])

	return goHorizontal(n, s, t, s, t, rowBound, colBound, 1, &min)
}

func goHorizontal(n *coupling.Node, s, t, u, v, rBound, cBound, pLen int, min *float64) bool {
	//log.Printf("entering node: (%v,%v)", u , v)
	var localmin float64
	edge := n.Adj[u][v]
	edge.To.Visited = true

	for i := 0; i < cBound; i++ {
		if i == v {
			// we are in node (u,v)
			continue
		}

		if n.Adj[u][i].To.Visited {
			// (u,i) has already been visited
			continue
		}

		if !(n.Adj[u][i].Basic) {
			// (u,i) is non-basic
			continue
		}

		// the node is basic

		// save the local minimun in case the next call is a dead-end
		localmin = *min

		// if next step is decrease update the global minimum
		*min = math.Min(*min, n.Adj[u][i].Prob)

		if goVertical(n, s, t, u, i, rBound, cBound, pLen+1, min) {
			// the path was finished
			updateEdge(edge, true, *min, n)
			//log.Printf("exiting node: (%v,%v)", u , v)
			return true
		}

		// restore global minimum if the path was a dead-end
		*min = localmin
	}

	edge.To.Visited = false
	//log.Printf("exiting node: (%v,%v)", u , v)
	return false
}

func goVertical(n *coupling.Node, s, t, u, v, rBound, cBound, pLen int, min *float64) bool {
	//log.Printf("entering node: (%v,%v)", u , v)
	var localmin float64
	edge := n.Adj[u][v]
	edge.To.Visited = true

	for i := 0; i < rBound; i++ {
		if i == u {
			// we are in node (u,v)
			continue
		}

		if i == s && v == t {
			if pLen%2 == 1 {
				log.Panic("stepping stone path cannot be uneven, since the intial node cannot be reached if the path is uneven")
			}
			//log.Printf("we have found node (%v,%v) again!", s, t)
			updateEdge(edge, false, *min, n)
			//log.Printf("exiting node: (%v,%v)", u , v)
			return true
		}

		if n.Adj[i][v].To.Visited {
			// (i,v) has already been visited
			continue
		}

		if !(n.Adj[i][v].Basic) {
			// (i,v) is non-basic
			continue
		}

		// the node is basic

		// save the local minimun in case the next call is a dead-end
		localmin = *min

		if goHorizontal(n, s, t, i, v, rBound, cBound, pLen+1, min) {
			updateEdge(edge, false, *min, n)
			//log.Printf("exiting node: (%v,%v)", u , v)
			return true
		}

		// restore global minimum if the path was a dead-end
		*min = localmin
	}

	edge.To.Visited = false
	//log.Printf("exiting node: (%v,%v)", u , v)
	return false
}

func updateEdge(edge *coupling.Edge, signal bool, min float64, n *coupling.Node) {
	if signal {
		log.Printf("increasing at cell (%v,%v)", edge.To.S, edge.To.T)
		// increase and set the node to basic
		edge.Prob += min
		
		if !edge.Basic {
			edge.Basic = true
			n.BasicCount++
		}
		
		// add the main node as a successor to the node if it not already is
		if !coupling.IsNodeInSlice(n, edge.To.Succ) {
			edge.To.Succ = append(edge.To.Succ, n)
		}

	} else {
		log.Printf("decreasing at cell (%v,%v)", edge.To.S, edge.To.T)
		edge.Prob -= min
		// if the line edge.Prob -= min, makes edge.Prob to zero, it is not a basic cell
		edge.Basic = !utils.ApproxEqual(edge.Prob, 0)
		
		// if no longer basic remove the main node as a successor for the node
		if !edge.Basic {
			coupling.DeleteNodeInSlice(n, &edge.To.Succ)
			n.BasicCount--
		}
	}

	edge.To.Visited = false
}
