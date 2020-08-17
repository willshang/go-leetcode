package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
}

// leetcode908_最小差值I
func smallestRangeI(A []int, K int) int {
	if len(A) == 1 {
		return 0
	}
	min := A[0]
	max := A[0]
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
		if A[i] < min {
			min = A[i]
		}
	}
	if max-min > 2*K {
		return max - min - 2*K
	}
	return 0
}
