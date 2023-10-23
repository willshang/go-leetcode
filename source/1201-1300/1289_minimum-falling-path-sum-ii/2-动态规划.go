package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(arr))
}

// leetcode1289_下降路径最小和II
func minFallingPathSum(arr [][]int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = arr[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
			for k := 0; k < n; k++ {
				if j != k {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+arr[i][j])
				}
			}
		}
	}
	res := math.MaxInt32
	for j := 0; j < n; j++ {
		res = min(res, dp[n-1][j])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
