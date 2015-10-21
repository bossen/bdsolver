package earthmover

import (
	"coupling"
	"math"
)

func SteppingStone(n *coupling.Node, s int, t int) {
	min, rowBound, colBound := 2.0, len(n.Adj), len(n.Adj[0])

	goHorizontal(n, s, t, s, t, rowBound, colBound, 1, &min)
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

		if approxFloatEqual(n.Adj[u][i].Prob, 0) {
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
			updateEdge(edge, true, *min)
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
			updateEdge(edge, false, *min)
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
			updateEdge(edge, false, *min)
			return true
		}

		// restore global minimum if the path was a dead-end
		*min = localmin
	}

	edge.To.Visited = false
	return false
}

func updateEdge(edge *coupling.Edge, signal bool, min float64) {
	if signal {
		// increase and set the node to basic
		edge.Prob += min
		edge.Basic = true

	} else {
		edge.Prob -= min
		// if the line edge.Prob -= min, makes edge.Prob to zero, it is not a basic cell
		edge.Basic = edge.Prob > 0
	}

	edge.To.Visited = false
}
