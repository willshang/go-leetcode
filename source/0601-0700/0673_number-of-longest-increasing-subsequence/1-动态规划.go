package main

import (
	"fmt"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
}

// leetcode673_最长递增子序列的个数
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 || nums == nil {
		return 0
	}
	dp := make([]int, n)
	count := make([]int, n)
	maxValue := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 {
					count[i] = count[i] + count[j]
				}
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		maxValue = max(maxValue, dp[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxValue {
			res = res + count[i]
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
