package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	// fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{-2, 0, -2}, 7))
}

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	i := 0
	for i < len(A)-1 && K > 0 {
		A[i] = -A[i]
		if A[i] > 0 && A[i] > A[i+1] {
			i++
		}
		K--
	}
	return sum(A)
}

func sum(A []int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		res = res + A[i]
	}
	return res
}
