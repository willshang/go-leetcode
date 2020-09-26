package main

import "fmt"

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}

// leetcode1014_最佳观光组合
func maxScoreSightseeingPair(A []int) int {
	res := 0
	maxValue := A[0] + 0
	// A[i]+A[j]+i-j=> max(A[i]+i)+(A[j]-j) (i<j)
	for j := 1; j < len(A); j++ {
		res = max(res, A[j]-j+maxValue)
		maxValue = max(maxValue, A[j]+j)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
