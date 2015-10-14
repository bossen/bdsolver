package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("hello world!")
}

func TestMy(t *testing.T) {
	t.Fail()
}
