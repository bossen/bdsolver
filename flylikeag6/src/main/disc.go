package main

import (
	"fmt"
)

func reachable(u int, v int, graph int) int {
  return graph
}

func transposegraph(graph int) int {
  return graph
}

func minus(a int, b int) int {
  return a
}

func union(a int, b int) int {
  return 1
}

func intersect(a int, b int) int {
  return a
}

func nextdemandedpair(w int, i int) (int, int) {
  return 1, 1
}

func setdistance(d *[256][256]int, u int, v int, value int) {
  d[u][v] = value
}

func reversegraphfunc(s int, t int, coupling int, nonzero *int, exact int, d [256][256]int) {
  pairs := intersect(reachable(s, t, coupling), exact)
  pairsize := 1 //len(pairs) TODO
  for i := 0; i < pairsize; i++ {
    u, v := nextdemandedpair(pairs, i)
    if d[u][v] > 0 {
      *nonzero = union(*nonzero, reachable(u, v, transposegraph(coupling)))
    }
  }
}

func setdistancetozero(s int, t int, nonzero int, exact *int, d *[256][256]int, coupling int) {
  pairs := minus(reachable(s, t, coupling), nonzero)
  pairsize := 1 //len(pairs) TODO
  for i := 0; i < pairsize; i++ {
    u, v := nextdemandedpair(pairs, i)
    setdistance(d, u, v, 0)
    *exact = 2
  }
}

func finda(coupling int, nonzero int) int {
  return 1
}

func findb(exact int, d[256][256]int, coupling int, nonzero int) int {
  return 1
}

func solvelinearsystem(lambda int, a int, b int) int {
  return 1
}

func getvalue(x int, u int, v int) int {
  return 1
}

func updatedistances(nonzero int, d *[256][256]int, x int) {
  pairsize := 1 //len(x) TODO
  for i := 0; i < pairsize; i++ {
    u, v := nextdemandedpair(x, i)
    setdistance(d, u, v, getvalue(x, u, v))
  }
}

func disc(lambda int, s int, t int, exact *int, coupling *int) {
  var _d [256][256] int
  d := &_d
  _d[1][1] = 1
  nonzero := 1
  reversegraphfunc(s, t, *coupling, &nonzero, *exact, *d)
  setdistancetozero(s, t, nonzero, exact, d, *coupling)
  a := finda(*coupling, nonzero)
  b := findb(*exact, *d, *coupling, nonzero)
  x := solvelinearsystem(lambda, a, b)
  updatedistances(nonzero, d, x)
  fmt.Println("discrepancy!")
}
