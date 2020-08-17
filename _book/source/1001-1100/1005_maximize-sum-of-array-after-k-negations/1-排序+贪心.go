package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	// fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{-2, 5, 0, 2, -2}, 3))
}

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	i := 0
	for i < len(A) && K > 0 {
		if A[i] < 0 {
			A[i] = -A[i]
			i++
			K--
		} else {
			break
		}
	}
	sort.Ints(A)
	if K%2 == 1 {
		A[0] = -A[0]
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
