package main

import "fmt"

func main() {
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}

// leetcode845_数组中的最长山脉
func longestMountain(A []int) int {
	n := len(A)
	left := make([]int, len(A))
	right := make([]int, len(A))
	for i := 1; i < n; i++ {
		if A[i-1] < A[i] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if A[i+1] < A[i] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 1; i < n-1; i++ {
		if left[i] > 0 && right[i] > 0 {
			res = max(res, left[i]+right[i]+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
