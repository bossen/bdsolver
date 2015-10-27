package earthmover

import (
    "markov"
    "coupling"
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


func randommatchingnew(m markov.MarkovChain, u int, v int) [][]*coupling.Edge {
    lenrow, lencol := 0, 0
    for i := range(m.Transitions[u]) {
        if m.Transitions[u][i] > 0 {
            lenrow += 1
        }
    }
    logger.Printf("earthmover.randommatching lenrow: %v", lenrow)

    for i := range(m.Transitions[v]) {
        if m.Transitions[v][i] > 0 {
            lencol += 1
        }
    }
    logger.Printf("earthmover.randommatching.lencol: %v", lencol)

    logger.Printf("earthmover.randommatching Making the matching matrix")
    matching := make([][]*coupling.Edge, lenrow, lenrow)
    for i := range(matching) {
        matching[i] = make([]*coupling.Edge, lencol, lencol)
    }

    return matching
}
