package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	}

	result := 0
	for x > 0 {
		temp := x % 10
		x = x / 10

		result = result*10 + temp
	}

	result = flag * result
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}
