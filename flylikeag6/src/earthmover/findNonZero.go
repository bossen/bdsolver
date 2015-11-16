package earthmover

import (
	"coupling"
	"log"
)

func findNonZero(s, t int, exact [][]bool, d [][]float64, c *coupling.Coupling) []*coupling.Node {
	n := coupling.FindNode(s, t, c)

	if n == nil {
		log.Panic("node chould not be found")
	}

	// finds all reachable from (s,t)
	r := coupling.Reachable(s, t, c)

	// remove nodes not exact or distance less than 0
	r = filterZeros(r, exact, d)

	// every node can always reach itself
	nonZeros := r

	// finds all reachable using the successor slice Succ
	for _, node := range r {
		nonZeros = findReverseReachable(node, nonZeros)
	}

	return nonZeros
}

func filterZeros(r []*coupling.Node, exact [][]bool, d [][]float64) []*coupling.Node {
	// copy the reachable set such that we remove elements, without removing them from the original
	temp := make([]*coupling.Node, len(r), len(r))
	copy(temp, r)

	for _, node := range temp {
		// if exact and distance above 0, skip it
		if exact[node.S][node.T] && d[node.S][node.T] > 0 {
			continue
		}
		// otherwise delete it
		deleteSucc(node, &r)
	}

	return r
}

func findReverseReachable(n *coupling.Node, r []*coupling.Node) []*coupling.Node {
	n.Visited = true

	for _, node := range n.Succ {
		if !succNode(node, r) {
			r = append(r, node)
		}

		if len(node.Succ) == 0 || node.Visited {
			continue
		}

		r = findReverseReachable(node, r)
	}

	n.Visited = false

	return r
}
