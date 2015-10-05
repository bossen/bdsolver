package main

import (
	"fmt"
)

func updatecoupling(coupling *int, w int, s int, t int) int {
	return 1
}

func union(visited *int, s int, t int) int {
	return 1
}

func matchcardinality(w int) int {
	return 0 
}

func nextdemandedpair(w int, i int) (int, int) {
	return 1, 1
}

func label(s int) int {
	return 1
}

func setdistance(d [256][256]int, u int, v int, n int) int {
	return 1
}

func notexact(u int, v int, exact *int) bool {
	return true
}

func randommatching(m int, u int, v int) int {
	return 1
}

func setpair(m int, s int, t int, w int, exact *int, visited *int, coupling *int) {
	fmt.Println("hi from setpair!")
	var d [256][256] int
	
	updatecoupling(coupling, w, s ,t)
	
	union(visited, s ,t)
	
	for i := 0; i < matchcardinality(w); i++ {
		u, v := nextdemandedpair(w, i)
		
		union(visited, u, v)
		
		if s == t {
			setdistance(d, u, v, 0)
		} else if label(s) == label(t) {
			setdistance(d, u, v, 1)
		} else if notexact(u, v, exact) {
			w2 := randommatching(m, u, v)
		    setpair(m, u, v, w2, exact, visited, coupling)
		}
	}
}
