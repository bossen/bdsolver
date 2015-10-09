package earthmover


import (
	"fmt"
    "sets"
)

func initexact() int {
	return 1
}

func initcoupling() int {
	return 1
}

func extractrandomfromset(tocompute *int) (int, int) {
	return 1, 1
}

func label(node int) string {
	return "red"
}


func removeedgesfromnodes(coupling *int, exact *int) int {
	return 1
}


func getoptimalschedule(d [256][256] int, u int, v int) int {
	return 1
}




func BipseudoMetric(n int) {
	var d [256][256]int
	visited := *sets.MakeMatrix(n)
	exact := initexact()
	coupling := initcoupling()
	tocompute := 0
	lambda := 1
	
	for !sets.EmptySet(tocompute) {
		s, t := extractrandomfromset(&tocompute)
		if label(s) != label(t) {
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

			for  { // TODO: add u, v in for loop
				// w := getoptimalschedule(d, u, v)
				w := 1
				setpair(1, s, t, w, &exact, &visited, &coupling)
				disc(lambda, s, t, &exact, &coupling)
			}
			exact = sets.Union(exact, reachable(s, t, coupling))
			removeedgesfromnodes(&coupling, &exact)
		}
		tocompute = sets.Intersect(tocompute, exact)


	}
	setpair(1, 1, 1, 1, &exact, &visited, &coupling)
	disc(1, 1, 1, &exact, &coupling)
	fmt.Println("hello world!")
	// return what?
}
