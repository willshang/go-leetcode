package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum(1, 2))
}

func maximum(a int, b int) int {
	// max(a,b) = (abs(a-b)+a+b)/2
	return (int(math.Abs(float64(a-b))) + a + b) / 2
}
