package main

import (
    "earthmover"
    "markov"
    "coupling"
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
    // fmt.Printf("%+v", mymarkov)
    
    testcoupling()
}

func testcoupling() {
    c := coupling.New()
    // c.Matchings[StatePair{1, 4}] = append(c.Matchings[StatePair{1,4}], (CouplingEdge{2,4, 1.0, 1}))
    c.Matchings[coupling.StatePair{1,4}] = append(c.Matchings[coupling.StatePair{1,4}], (coupling.CouplingEdge{2, 1, 1/3, 0}))
    c.Matchings[coupling.StatePair{1,4}] = append(c.Matchings[coupling.StatePair{1,4}], (coupling.CouplingEdge{3, 2, 1/3, 0}))
    c.Matchings[coupling.StatePair{1,4}] = append(c.Matchings[coupling.StatePair{1,4}], (coupling.CouplingEdge{4, 3, 1/6, 0}))
    c.Matchings[coupling.StatePair{1,4}] = append(c.Matchings[coupling.StatePair{1,4}], (coupling.CouplingEdge{6, 3, 1/6, 0}))

    c.Matchings[coupling.StatePair{3,4}] = append(c.Matchings[coupling.StatePair{3, 4}], (coupling.CouplingEdge{2, 1, 1/3, 0}))
    c.Matchings[coupling.StatePair{3,4}] = append(c.Matchings[coupling.StatePair{3, 4}], (coupling.CouplingEdge{2, 2, 1/6, 0}))
    c.Matchings[coupling.StatePair{3,4}] = append(c.Matchings[coupling.StatePair{3, 4}], (coupling.CouplingEdge{3, 2, 1/6, 0}))
    c.Matchings[coupling.StatePair{3,4}] = append(c.Matchings[coupling.StatePair{3, 4}], (coupling.CouplingEdge{3, 3, 1/6, 0}))
    coupling.Reachable(1, 4, c)
    // fmt.Printf("%+v", c)

}
