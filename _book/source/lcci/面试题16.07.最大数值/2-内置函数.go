package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum(1, 2))
}

func maximum(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
