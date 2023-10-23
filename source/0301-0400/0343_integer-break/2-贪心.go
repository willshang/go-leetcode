package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(integerBreak(10))
}

// leetcode343_整数拆分
func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	timesOf3 := n / 3
	if n-timesOf3*3 == 1 {
		timesOf3 = timesOf3 - 1
	}
	timesOf2 := (n - timesOf3*3) / 2
	return int(math.Pow(float64(2), float64(timesOf2))) *
		int(math.Pow(float64(3), float64(timesOf3)))
}
