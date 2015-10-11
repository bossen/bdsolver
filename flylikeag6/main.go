package main

import (
    "earthmover"
)

func main() {
	earthmover.BipseudoMetric(32)
    mymarkov := markov.New()
    fmt.Printf("%+v", mymarkov)
}
