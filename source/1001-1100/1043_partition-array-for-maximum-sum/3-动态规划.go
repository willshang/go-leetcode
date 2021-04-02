package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7}, 3))
}

// leetcode1043_分隔数组以得到最大和
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		maxValue := arr[i]
		for j := i; j >= 0 && j > i-k; j-- {
			maxValue = max(maxValue, arr[j])
			if j > 0 {
				dp[i] = max(dp[i], dp[j-1]+maxValue*(i-j+1))
			} else {
				dp[i] = max(dp[i], maxValue*(i+1))
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
