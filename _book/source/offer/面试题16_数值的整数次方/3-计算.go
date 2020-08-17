package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(float64(2.0), 10))
}

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
