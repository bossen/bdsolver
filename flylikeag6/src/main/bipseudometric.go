package main


import (
	"fmt"
)

func initvisited() int {
	return 1
}

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

func union(a int, b int, c int) int{
	return a + b
}

func reachable(s int, t int) int{
	return s + t
}

func removeedgesfromnodes(coupling *int, exact *int) int {
	return 1
}


func getoptimalschedule(d [256][256] int, u int, v int) int {
	return 1
}

var emptysetnotalways = 2

func emptyset(set int) bool {
	emptysetnotalways -= 1
	if emptysetnotalways < 0 {
		return true
	} else {
		return false
	}
}


func intersect(a int, b int) int {
	return 1
}

func bipseudometric() {
	var d [256][256]int
	visited := initvisited()
	exact := initexact()
	coupling := initcoupling()
	tocompute := 0
	lambda := 1
	
	for !emptyset(tocompute) {
		s, t := extractrandomfromset(&tocompute)
		if label(s) == label(t) {
			d[s][t] = 1
			exact = union(exact, s, t)
			visited = union(visited, s, t)
		} else if s == t {
			d[s][t] = 0
			exact = union(exact, s, t)
			visited = union(visited, s, t)

		} else {
			// if s,t not in visited ...
			
			disc(lambda, s, t, &exact, &coupling)

			for  { // TODO: add u, v in for loop
				// w := getoptimalschedule(d, u, v)
				w := 1
				setpair(1, s, t, w, &exact, &visited, &coupling)
				disc(lambda, s, t, &exact, &coupling)
			}
			exact = union(exact, reachable(s, t), 0)
			removeedgesfromnodes(&coupling, &exact)
		}
		tocompute = intersect(tocompute, exact)


	}
	setpair(1, 1, 1, 1, &exact, &visited, &coupling)
	disc(1, 1, 1, &exact, &coupling)
	fmt.Println("hello world!")
	// return what?
}
