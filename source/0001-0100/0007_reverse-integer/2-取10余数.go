package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(120))
}

// leetcode 7 整数反转
func reverse(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		result = result*10 + temp
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return result
}
