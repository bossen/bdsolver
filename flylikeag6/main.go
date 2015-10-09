package main

import (
    "earthmover"
    "markov"
    "fmt"
)

func main() {
	earthmover.BipseudoMetric(32)
    mymarkov := markov.New()
    fmt.Printf("%+v", mymarkov)
}
