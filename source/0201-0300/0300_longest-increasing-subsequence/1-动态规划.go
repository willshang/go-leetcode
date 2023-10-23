package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

// leetcode300_最长上升子序列
/*
dp[i] = max(dp[j]+1, dp[i]),其中0<=j<i, nums[j]<nums[i])
*/
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
