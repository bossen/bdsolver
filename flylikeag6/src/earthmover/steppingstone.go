package earthmover

import (
	"coupling"
)

func SteppingStone(n *coupling.Node, s int, t int) {
	min, rowBound, colBound := 2.0, len(n.Adj), len(n.Adj[0])
	
	goHorizontal(n, s, t, s, t, rowBound, colBound, 1, true, &min)
}

func goHorizontal(n *coupling.Node, s, t, u, v, rBound, cBound, pLen int, signal bool, min *float64) bool {
	var localmin float64
	edge := n.Adj[u][v]
	edge.To.Color = 1
	
	for i:= 0; i < cBound; i++ {
		if i == v {
			// we are in node (u,v)
			continue
		}
		
		if u == s && i == t && pLen > 3 {
			// we have finished the path
			updateEdge(edge, signal, *min)
			return true	
		} 
		
		if n.Adj[u][i].To.Color == 1 {
			// (u,i) has already been visited
			continue
		}
		
		if !(n.Adj[u][i].IsBasic) {
			// (u,i) is non-basic
			continue
		}
		
		// the node is basic
		
		if approxFloatEqual(n.Adj[u][i].Prob, 0) && signal {
			// if next step is decrease and the prob at (u,i) is 0 we
			// cannot go there
			continue
		}
		
		// save the local minimun in case the next call is a dead-end
		localmin = *min
		
		if signal {
			// if next step is decrease update the global minimum
			min = findMin(*min, n.Adj[u][i].Prob)
		}
		
		if goVertical(n, s, t, u, i, rBound, cBound, pLen + 1, !signal, min) {
			// the path was finished
			updateEdge(edge, signal, *min)
			return true
		}
		
		// restore global minimum if the path was a dead-end
		min = &localmin
	}
	
	edge.To.Color = 0
	return false
}

func goVertical(n *coupling.Node, s, t, u, v, rBound, cBound, pLen int, signal bool, min *float64) bool {
	var localmin float64
	edge := n.Adj[u][v]
	edge.To.Color = 1
	
	for i:= 0; i < rBound; i++ {
		if i == u {
			// we are in node (u,v)
			continue
		}
		
		if i == s && v == t && pLen > 3 {
			// we have finished the path
			updateEdge(edge, signal, *min)
			return true	
		} 
		
		if n.Adj[i][v].To.Color == 1 {
			// (i,v) has already been visited
			continue
		}
		
		if !(n.Adj[i][v].IsBasic) {
			// (i,v) is non-basic
			continue
		}
		
		// the node is basic
		
		if approxFloatEqual(n.Adj[i][v].Prob, 0) && signal {
			// if next step is decrease and the prob at (i,v) is 0 we
			// cannot go there
			continue
		}
		
		// save the local minimun in case the next call is a dead-end
		localmin = *min
		
		if signal {
			// if next step is decrease update the global minimum
			min = findMin(*min, n.Adj[i][v].Prob)
		}
		
		if goHorizontal(n, s, t, i, v, rBound, cBound, pLen + 1, !signal, min) {
			updateEdge(edge, signal, *min)
			return true
		}
		
		// restore global minimum if the path was a dead-end
		min = &localmin
	}
	
	return false	
}

func updateEdge(edge *coupling.Edge, signal bool, min float64) {
	if signal {
		// increase and set the node to basic
		edge.Prob = edge.Prob + min
		edge.IsBasic = true
		
	} else {
		// decrase and set the node to non-basic if not bigger than 0
		edge.Prob = edge.Prob - min
				
		if edge.Prob > 0 {
			edge.IsBasic = true
		} else {
			edge.IsBasic = false
		}
	}
	
	edge.To.Color = 0
}

func findMin(a float64, b float64) *float64 {
	if a < b {
		return &a
	} else {
		return &b
	}
}
