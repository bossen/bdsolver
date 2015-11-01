package earthmover

import (
    "markov"
    "coupling"
    "log"
)

func randommatching(m markov.MarkovChain, u int, v int) [][]float64 {
	j, k, n := 0, 0, len(m.Labels)
	uTransitions := make([]float64, n, n)
	vTransitions := make([]float64, n, n)

	copy(uTransitions, m.Transitions[u])
	copy(vTransitions, m.Transitions[v])

	matching := make([][]float64, n, n)

	for i := range matching {
		matching[i] = make([]float64, n, n)
	}

	for j < n && k < n {
		if approxFloatEqual(uTransitions[j], vTransitions[k]) {
			matching[j][k] = uTransitions[j]
			j++
			k++
		} else if uTransitions[j] < vTransitions[k] {
			matching[j][k] = uTransitions[j]
			vTransitions[k] = vTransitions[k] - uTransitions[j]
			j++
		} else {
			matching[j][k] = vTransitions[k]
			uTransitions[j] = uTransitions[j] - vTransitions[k]
			k++
		}
	}

	return matching
}


func randommatchingnew(m markov.MarkovChain, u int, v int, c *coupling.Coupling) *coupling.Node {
    n := len(m.Transitions[u])
    lenrow, lencol := 0, 0
    l, k := 0, 0
    
    log.Printf("copy the transitions from the states %v and %v", u, v)
    uTransitions := make([]float64, n, n)
	vTransitions := make([]float64, n, n)

	copy(uTransitions, m.Transitions[u])
	copy(vTransitions, m.Transitions[v])
	
	log.Println(uTransitions)
	log.Println(vTransitions)
    
    for i := 0; i < n; i++ {
        if m.Transitions[u][i] > 0 {
            lenrow += 1
        }
        if m.Transitions[v][i] > 0 {
			lencol += 1
		}
    }
    log.Printf("row and column length are: %v and %v", lenrow, lencol)
    
    rowindex := make([]int, lenrow, lenrow)
    colindex := make([]int, lencol, lencol)
    
    for i := 0; i < n; i++ {
		if m.Transitions[u][i] > 0 {
            rowindex[l] = i
            l++
        }
        if m.Transitions[v][i] > 0 {
			colindex[k] = i
            k++
		}
	}
	log.Printf("row index: %s", rowindex)
	log.Printf("column index: %s", colindex)

    log.Printf("earthmover.randommatching Making the matching matrix")
    matching := make([][]*coupling.Edge, lenrow, lenrow)
    
    for i := range(matching) {
        matching[i] = make([]*coupling.Edge, lencol, lencol)
    }
    
    for i := 0; i < lenrow; i++ {
		for j := 0; j < lencol; j++ {
			var node *coupling.Node
			
			if rowindex[i] <= rowindex[j] {
				node = coupling.FindNode(rowindex[i], colindex[j], c)
			} else {
				node = coupling.FindNode(colindex[j], rowindex[i], c)
			}
			
			if node == nil {
				node = &coupling.Node{S: rowindex[i], T: colindex[j]}
				c.Nodes = append(c.Nodes, node)
			}
			
			matching[i][j] = &coupling.Edge{To: node}
		}
	}
	
	l, k = 0, 0
	
	for l < lenrow && k < lencol {
		if approxFloatEqual(uTransitions[rowindex[l]], vTransitions[colindex[k]]) {
			matching[l][k].Prob = uTransitions[rowindex[l]]
			
			if !(l + 1 == lenrow && k + 1 == lencol) {
				matching[l][k+1].Basic = true
			}
			
			matching[l][k].Basic = true
			
			l++
			k++
		} else if uTransitions[rowindex[l]] < vTransitions[colindex[k]] {
			matching[l][k].Prob = uTransitions[rowindex[l]]
			vTransitions[colindex[k]] = vTransitions[colindex[k]] - uTransitions[rowindex[l]]
			
			matching[l][k].Basic = true
			
			l++
		} else {
			matching[l][k].Prob = vTransitions[colindex[k]]
			uTransitions[rowindex[l]] =  uTransitions[rowindex[l]] - vTransitions[colindex[k]]
			
			matching[l][k].Basic = true
			
			k++
		}
	}
	
	node := coupling.FindNode(u, v, c)
	
	if node == nil {
		node = &coupling.Node{S: u, T: v}
		c.Nodes = append(c.Nodes, node)
	}
	
	for i := 0; i < lenrow; i++ {
		for j := 0; j < lencol; j++ {
			log.Println("At: %v and %v", i, j)
			log.Println(matching[i][j].Basic)
		}
	}
	
	node.Adj = matching

    return node
}
