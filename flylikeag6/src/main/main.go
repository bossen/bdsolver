package main

import (
	"fmt"
)

func main() {
	visited := 1
	exact := 1
	coupling := 1

	fmt.Println("hello world!")
	setpair(1, 1, 1, 1, &exact, &visited, &coupling)
	disc(1, 1, 1, &exact, &coupling)
}
