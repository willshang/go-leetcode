package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(7, -3))
}

func divide(dividend int, divisor int) int {
	res := dividend / divisor
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
