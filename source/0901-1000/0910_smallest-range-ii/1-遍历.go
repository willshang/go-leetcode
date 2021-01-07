package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(smallestRangeII([]int{0, 10}, 2))
	fmt.Println(smallestRangeII([]int{1, 3, 6}, 3))
}

// leetcode910_最小差值II
func smallestRangeII(A []int, K int) int {
	sort.Ints(A)
	n := len(A)
	res := A[n-1] - A[0]
	// 排序后，为了最小差值，必定是A[0,i]+K,A[i+1:]-K
	// 最小值落在A[0]+K，A[i+1]-K之中
	// 最大值落在A[n-1]-K，A[i]+K之中
	for i := 0; i < n-1; i++ {
		minValue := min(A[0]+K, A[i+1]-K)
		maxValue := max(A[n-1]-K, A[i]+K)
		res = min(maxValue-minValue, res)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
