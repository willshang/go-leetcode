package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(integerReplacement(8))
}

// leetcode397_整数替换
func integerReplacement(n int) int {
	if n < 3 {
		return n - 1
	}
	if n == math.MinInt32 {
		return 32
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	}
	a := integerReplacement(n+1) + 1
	b := integerReplacement(n-1) + 1
	if a > b {
		return b
	}
	return a
}
