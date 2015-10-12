package earthmover


import (
	"fmt"
    "sets"
    "markov"
)

func initexact() int {
	return 1
}

func initcoupling() int {
	return 1
}

func extractrandomfromset(tocompute *[][]bool) (int, int) {
  for i := range *tocompute {
    for j:= range *tocompute {
      if (*tocompute)[i][j] == true {
        return i, j
      }
    }
  }
  panic("Tried to extract random element from empty set!")
}

func removeedgesfromnodes(coupling *int, exact *int) int {
	return 1
}


func getoptimalschedule(d [256][256] int, u int, v int) int {
	return 1
}

func isOptimal() bool {
  return true
}

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

// credits to https://gist.github.com/cevaris/bc331cbe970b03816c6b
func approxFloatEqual(a, b float64) bool {
	var epsilon float64 = 0.00000001
	
	if ((a - b) < epsilon && (b - a) < epsilon) {
		return true
	}
	return false
}

func BipseudoMetric(m markov.MarkovChain,  lambda int, tocompute *[][]bool) {
	var d [256][256]int
    n := len(m.Transitions)
	visited := *sets.MakeMatrix(n)
	exact := initexact()
	coupling := initcoupling()
	w2 := randommatching(m, 0, 1)
	fmt.Println(w2)
	
	for !sets.EmptySet(tocompute) {
		s, t := extractrandomfromset(tocompute)
    fmt.Println(s)
    fmt.Println(t)
		if m.Labels[s] != m.Labels[t] {
			d[s][t] = 1
			exact = sets.UnionNode(exact, s, t)
			visited[s][t] = true
		} else if s == t {
			d[s][t] = 0
			exact = sets.UnionNode(exact, s, t)
			visited[s][t] = true

		} else {
			// if s,t not in visited ...

            disc(lambda, s, t, &exact, &coupling)

			for !isOptimal()  { // TODO: add u, v in for loop. While loop check for NOT optimal matching
				// w := getoptimalschedule(d, u, v)
				w := randommatching(m, s ,t)
				setpair(m, s, t, w, &exact, &visited, &coupling)
				disc(lambda, s, t, &exact, &coupling)
			}
			exact = sets.Union(exact, reachable(s, t, coupling))
			removeedgesfromnodes(&coupling, &exact)
		}
    
		tocompute = sets.IntersectReal(*tocompute, *tocompute) //TODO THIS IS WRONG, use exact as second parameter, instead of tocompute twice

	}
	setpair(m, 1, 1, w2, &exact, &visited, &coupling)
	disc(1, 1, 1, &exact, &coupling)
	fmt.Println("hello world!")
	// return what?
}
