package main

import (
    "earthmover"
    "fmt"
    "markov"
)

func main() {
	earthmover.BipseudoMetric(32)
    mymarkov := markov.New()
    fmt.Printf("%+v", mymarkov)
}
