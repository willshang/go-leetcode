package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	// fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	//fmt.Println(largestSumAfterKNegations([]int{-2, 0, -2}, 7))
}

func largestSumAfterKNegations(A []int, K int) int {
	for K > 0 {
		minIndex, minValue := findMin(A)
		if minValue > 0 {
			break
		}
		A[minIndex] = -A[minIndex]
		K--
	}
	if K%2 == 1 {
		minIndex, _ := findMin(A)
		A[minIndex] = -A[minIndex]
	}
	res := 0
	for i := 0; i < len(A); i++ {
		res = res + A[i]
	}
	return res
}

func findMin(A []int) (int, int) {
	res := A[0]
	index := 0
	for i := 1; i < len(A); i++ {
		if res > A[i] {
			res = A[i]
			index = i
		}
	}
	return index, res
}
