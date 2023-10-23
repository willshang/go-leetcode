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
	dp := make([]int, m+1)
	res := math.MinInt32
	for i := 1; i <= n; i++ {
		for j := m; j >= 1; j-- {
			if A[i-1] == B[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0 // 需要清0
			}
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res
}
