package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(2))
}

// leetcode633_三个数的最大乘积
func judgeSquareSum(c int) bool {
	for i := 0; i <= int(math.Sqrt(float64(c))); i++ {
		b := c - i*i
		s := int(math.Sqrt(float64(b)))
		if s*s == b {
			return true
		}
	}
	return false
}
