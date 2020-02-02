package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(1))
}

func mySqrt(x int) int {
	result := math.Floor(math.Sqrt(float64(x)))
	return int(result)
}
