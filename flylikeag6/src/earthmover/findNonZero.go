package earthmover

import (
	"coupling"
)

func findNonZero(n *coupling.Node, exact [][]bool, d [][]float64, c *coupling.Coupling) []*coupling.Node {
	// finds all reachable from (s,t)
	reachables := coupling.Reachable(n)

	// remove nodes not exact or distance less than 0
	reachablesNonZero := filterZeros(reachables, exact, d)

	// every node can always reach itself, so we copy the reachablesNonZeros into the final nonZeros slice
	nonZeros := make([]*coupling.Node, len(reachablesNonZero))
	copy(nonZeros, reachablesNonZero)

	// finds all reachable using the successor slice Succ
	for _, node := range reachablesNonZero {
		nonZeros = findReverseReachable(node, nonZeros)
	}

	return nonZeros
}

func filterZeros(reachables []*coupling.Node, exact [][]bool, d [][]float64) []*coupling.Node {
	// copy the reachable set such that we remove elements, without removing them from the original
	reachablesCopy := make([]*coupling.Node, len(reachables), len(reachables))
	copy(reachablesCopy, reachables)

	for _, node := range reachablesCopy {
		// if exact and distance above 0, skip it
		if exact[node.S][node.T] && d[node.S][node.T] > 0 {
			continue
		}
		// otherwise delete it
		coupling.DeleteNodeInSlice(node, &reachables)
	}

	return reachables
}

func findReverseReachable(node *coupling.Node, reachables []*coupling.Node) []*coupling.Node {
	node.Visited = true

	for _, succ := range node.Succ {
		if !coupling.IsNodeInSlice(succ, reachables) {
			reachables = append(reachables, succ)
		}

		if len(succ.Succ) == 0 || succ.Visited {
			continue
		}

		reachables = findReverseReachable(succ, reachables)
	}

	node.Visited = false

	return reachables
}
