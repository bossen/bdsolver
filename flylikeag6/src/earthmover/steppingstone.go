package earthmover

import (
	"coupling"
)

func SteppingStone(n *coupling.Node, s int, t int) {
	rowBound, colbound := len(*n.Adj), len((*n.Adj)[0])
	var done bool
	
	for i := s + 1; i < rowBound, i++ {
		
	}
}

func goHorizontal(n *coupling.Node, s, t, u, v, rBound, cBound, pLen, signal bool, min *float64) bool {
	var localmin float64
	edge := (*n.Adj)[u][v]
	
	for i:= 0; i < cBound; i++ {
		if i == v {
			continue
		}
		
		if u == s && i == t && pLen > 3 {
			updateEdge(&edge, signal, min)
			return true
		} else if !((*n.Adj)[u][i].IsBasic) {
			continue
		} 
		
		// the node is basic
		
		if !signal {
			localmin = min
			min = findMin(min, edge.Prob)
		}
		
		if approxFloatEqual((*n.Adj)[u][i].Prob, 0) && !signal {
			continue
		}
		
		if goVertical(&n, s, t, u, i, rBound, cBound, pLen + 1, !signal, &min) {
			updateEdge(&edge, signal, min)
			return true
		}
		
		min = localmin
	}
	
	return false
}

func goVertical(n *coupling.Node, s, t, u, v, rBound, cBound, pLen, signal int, min float64) bool {
	var localmin float64
	edge := (*n.Adj)[u][v]
	
	for i:= 0; i < rBound; i++ {
		if i == u {
			continue
		}
		
		if i == s && v == t && pLen > 3 {
			updateEdge(&edge, signal, min)
			return true	
		} else if !((*n.Adj)[i][v].IsBasic) {
			continue
		} 
		
		// the node is basic
		
		if !signal {
			localmin = min
			min = findMin(min, edge.Prob)
		}
		
		if approxFloatEqual((*n.Adj)[i][v].Prob, 0) && !signal {
			continue
		}
		
		if goVertical(&n, s, t, i, v, rBound, cBound, pLen + 1, !signal, &min) {
			updateEdge(&edge, signal, min)
			return true
		}
		
		min = localmin
	}
	
	return false	
}

func updateEdge(edge *coupling.Edge, signal int, min float64) {
	if signal == 1 {
		edge.Prob = edge.Prob + min
	} else {
		min = findMin(min, edge.Prob)
		edge.Prob = edge.Prob - min
				
		if approxFloatEqual(edge.Prob, 0) {
			edge.IsBasic = false
		}
	}
}

func findMin(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}
