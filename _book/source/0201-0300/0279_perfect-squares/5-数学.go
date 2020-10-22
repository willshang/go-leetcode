package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	if judge(n) {
		return 1
	}
	res := n
	for res%4 == 0 {
		res = res / 4
	}
	if res%8 == 7 {
		return 4
	}
	for i := 1; i*i < n; i++ {
		if judge(n - i*i) {
			return 2
		}
	}
	return 3
}

func judge(n int) bool {
	value := int(math.Sqrt(float64(n)))
	return value*value == n
}
