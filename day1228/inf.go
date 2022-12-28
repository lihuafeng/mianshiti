package main

import (
	"fmt"
	"math"
)

func main() {
	// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
	x := math.Inf(2)
	y := math.Inf(-2)
	fmt.Print(y)
	switch {
	case x < 0, x > 0:
		fmt.Println(x)
	case x == 0:
		fmt.Println("zero")
	default:
		fmt.Println("something else")
	}
}
