package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7}, 3))
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		start := max(0, i-k)
		maxValue := math.MinInt32
		for j := i - 1; j >= start; j-- {
			maxValue = max(maxValue, arr[j])
			dp[i] = max(dp[i], dp[j]+maxValue*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
