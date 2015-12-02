package earthmover

import (
	"coupling"
	"log"
	"markov"
    "utils"
)

func matchingDimensions(u, v []float64, n int) (int, int) {
	lenrow, lencol := 0, 0
	
	for i:= 0; i < n; i++ {
		if u[i] > 0 {
			lenrow += 1
		}
		if v[i] > 0 {
			lencol += 1
		}
	}
	
	return lenrow, lencol
}

func setMatchingIndexes(u, v []float64, n int, row, col []int) {
	l, k := 0, 0
	for i := 0; i < n; i++ {
		if u[i] > 0 {
			row[l] = i
			l++
		}
		if v[i] > 0 {
			col[k] = i
			k++
		}
	}
}

func filloutAdj(row, col []int, lenrow, lencol int, w [][]*coupling.Edge, c *coupling.Coupling) {
	for i := 0; i < lenrow; i++ {
		for j := 0; j < lencol; j++ {
			var node *coupling.Node
			s, t := utils.GetMinMax(row[i], col[j])

			node = coupling.FindNode(s, t, c)

			if node == nil {
				node = &coupling.Node{S: s, T: t, Succ: []*coupling.Node{}}
				c.Nodes = append(c.Nodes, node)
			}

			w[i][j] = &coupling.Edge{To: node}
		}
	}
}

func findFeasibleMatching(m markov.MarkovChain, u int, v int, c *coupling.Coupling) *coupling.Node {
	n := len(m.Transitions[u])

	u, v = utils.GetMinMax(u, v)
	
	// tries to find the node (u,v) in c, if not make a new one and add to c
	node := coupling.FindNode(u, v, c)

	if node == nil {
		node = &coupling.Node{S: u, T: v, Succ: []*coupling.Node{}}
		c.Nodes = append(c.Nodes, node)
	}

	log.Printf("copy the transitions from the states %v and %v", u, v)
	uTransitions := make([]float64, n, n)
	vTransitions := make([]float64, n, n)

	// copies the transitions of u and v such that we do not makes changes in the markov chain
	copy(uTransitions, m.Transitions[u])
	copy(vTransitions, m.Transitions[v])

	log.Printf("Transitions for %v: %s", u, uTransitions)
	log.Printf("Transitions for %v: %s", v, vTransitions)
	
	// finds the length of the rows and columns in the matching for u and v
	lenrow, lencol := matchingDimensions(uTransitions, vTransitions, n)
	
	log.Printf("row and column length are: %v and %v", lenrow, lencol)

	rowindex := make([]int, lenrow, lenrow)
	colindex := make([]int, lencol, lencol)
	
	// finds the row and column indexs for the u and v matching
	setMatchingIndexes(uTransitions, vTransitions, n, rowindex, colindex)
	
	log.Printf("row index: %s", rowindex)
	log.Printf("column index: %s", colindex)

	matching := make([][]*coupling.Edge, lenrow, lenrow)

	for i := range matching {
		matching[i] = make([]*coupling.Edge, lencol, lencol)
	}
	
	// fills out the matching with node pointers, using nodes already in the coupling c if they exist
	filloutAdj(rowindex, colindex, lenrow, lencol, matching, c)

	// completes the matching by inserting probabilities and setting appropriate cells to basic
	i, j := 0, 0
	for i < lenrow && j < lencol {
		if utils.ApproxEqual(uTransitions[rowindex[i]], vTransitions[colindex[j]]) {
			matching[i][j].Prob = uTransitions[rowindex[i]]

			// check if we are in the lower right corner, such that we do not get an out of bounds error
			if !(i+1 == lenrow && j+1 == lencol) {
				matching[i][j+1].Basic = true
				node.BasicCount++
			}

			matching[i][j].Basic = true
			matching[i][j].To.Succ = append(matching[i][j].To.Succ, node)
			node.BasicCount++

			i++
			j++
		} else if uTransitions[rowindex[i]] < vTransitions[colindex[j]] {
			matching[i][j].Prob = uTransitions[rowindex[i]]
			vTransitions[colindex[j]] = vTransitions[colindex[j]] - uTransitions[rowindex[i]]

			matching[i][j].Basic = true
			matching[i][j].To.Succ = append(matching[i][j].To.Succ, node)
			node.BasicCount++

			i++
		} else {
			matching[i][j].Prob = vTransitions[colindex[j]]
			uTransitions[rowindex[i]] = uTransitions[rowindex[i]] - vTransitions[colindex[j]]

			matching[i][j].Basic = true
			matching[i][j].To.Succ = append(matching[i][j].To.Succ, node)
			node.BasicCount++

			j++
		}
	}

    utils.LogMatching(matching, lenrow, lencol)

	node.Adj = matching

	return node
}
