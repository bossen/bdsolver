package earthmover

import (
	"coupling"
	"math"
    "utils"
)

func SteppingStone(n *coupling.Node, s int, t int) bool {
	min, rowBound, colBound := 2.0, len(n.Adj), len(n.Adj[0])

	return goHorizontal(n, s, t, s, t, rowBound, colBound, 1, &min)
}

func goHorizontal(n *coupling.Node, s, t, u, v, rBound, cBound, pLen int, min *float64) bool {
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

		if utils.ApproxEqual(n.Adj[u][i].Prob, 0) {
			// if next step is decrease and the prob at (u,i) is 0 we
			// cannot go there
			continue
		}

		// save the local minimun in case the next call is a dead-end
		localmin = *min

		// if next step is decrease update the global minimum
		*min = math.Min(*min, n.Adj[u][i].Prob)

		if goVertical(n, s, t, u, i, rBound, cBound, pLen+1, min) {
			// the path was finished
			updateEdge(edge, true, *min, n)
			return true
		}

		// restore global minimum if the path was a dead-end
		*min = localmin
	}

	edge.To.Visited = false
	return false
}

func goVertical(n *coupling.Node, s, t, u, v, rBound, cBound, pLen int, min *float64) bool {
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
				panic("stepping stone path cannot be uneven, since the intial node cannot be reached if the path is uneven")
			}
			// we have finished the path
			updateEdge(edge, false, *min, n)
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
			return true
		}

		// restore global minimum if the path was a dead-end
		*min = localmin
	}

	edge.To.Visited = false
	return false
}

func updateEdge(edge *coupling.Edge, signal bool, min float64, n *coupling.Node) {
	if signal {
		// increase and set the node to basic
		edge.Prob += min
		edge.Basic = true
		
		// add the main node as a successor to the node if it not already is
		if !succNode(n, edge.To.Succ) {
			edge.To.Succ = append(edge.To.Succ, n)
		}

	} else {
		edge.Prob -= min
		// if the line edge.Prob -= min, makes edge.Prob to zero, it is not a basic cell
		edge.Basic = edge.Prob > 0
		
		// if no longer basic remove the main node as a successor for the node
		if !edge.Basic {
			deleteSucc(n, &edge.To.Succ)
		}
	}

	edge.To.Visited = false
}

func succNode(n *coupling.Node, succ []*coupling.Node) bool {
	for _, i := range succ {
		if i == n {
			return true
		}
	}
	return false
}

func deleteSucc(n *coupling.Node, succ *[]*coupling.Node) {
	for i := 0; i < len(*succ); i++ {
		if (*succ)[i] == n {
			// https://github.com/golang/go/wiki/SliceTricks
			(*succ)[i] = (*succ)[len(*succ)-1]
			(*succ)[len(*succ)-1] = nil
			(*succ) = (*succ)[:len(*succ)-1]
			break
		}
	}
	return
}
