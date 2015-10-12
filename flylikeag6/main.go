package main

import (
    "earthmover"
    "fmt"
    "markov"
    "sets"
)


//Computes all the possible combinations of the different nodes. This could be optimized, by setting everything below the i == j diagonal to false.
func initToCompute(n int) *[][]bool {
  toCompute := *sets.MakeMatrix(n)
  for i := range toCompute {
    for j := range toCompute {
      if i == j {
        toCompute[i][j] = false
      } else {
        toCompute[i][j] = true
      }
    }
  }
  return &toCompute
}
func main() {
    mymarkov := markov.New()
	tocompute := initToCompute(len(mymarkov.Transitions))
	earthmover.BipseudoMetric(mymarkov, 32, tocompute)
    fmt.Printf("%+v", mymarkov)
}
