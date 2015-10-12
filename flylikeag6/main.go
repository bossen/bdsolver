package main

import (
    "earthmover"
    "fmt"
    "markov"
    "coupling"
)

func main() {
	earthmover.BipseudoMetric(32)
    mymarkov := markov.New()
    fmt.Printf("%+v", mymarkov)
    
    n1 := coupling.StatePair{0, 2}
    n2 := coupling.StatePair{0, 2}
    fmt.Println(n1 == n2)
}
