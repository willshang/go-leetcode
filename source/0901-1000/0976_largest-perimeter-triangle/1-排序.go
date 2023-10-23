package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
}

// leetcode976_三角形的最大周长
func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
