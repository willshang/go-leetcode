package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	fmt.Println(findLength([]int{1, 0, 0, 0, 1}, []int{1, 0, 0, 1, 1}))
}

func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	res := math.MinInt32
	for i := 0; i < n; i++ {
		length := min(n-i, m)
		maxLength := getMaxLength(A, B, i, 0, length)
		res = max(res, maxLength)
	}
	for i := 0; i < m; i++ {
		length := min(n, m-i)
		maxLength := getMaxLength(A, B, 0, i, length)
		res = max(res, maxLength)
	}
	return res
}

func getMaxLength(A, B []int, a, b int, length int) int {
	res := 0
	count := 0
	for i := 0; i < length; i++ {
		if A[a+i] == B[b+i] {
			count++
		} else {
			count = 0
		}
		res = max(res, count)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
