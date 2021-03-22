package main

import "fmt"

func main() {
	fmt.Println(getMaxLen([]int{-82, 62, 54, -63,
		-85, 53, -60, -59,
		29, 32, 59, -54,
		-29, -45}))
}

// leetcode1567_乘积为正数的最长子数组长度
func getMaxLen(nums []int) int {
	dp := make([][2]int, len(nums)+1)
	res := 0
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			dp[i][0] = 0
			dp[i][1] = 0
		} else if nums[i-1] > 0 {
			dp[i][0] = dp[i-1][0] + 1
			if dp[i-1][1] != 0 {
				dp[i][1] = dp[i-1][1] + 1
			} else {
				dp[i][1] = 0
			}
		} else {
			if dp[i-1][1] != 0 {
				dp[i][0] = dp[i-1][1] + 1
			} else {
				dp[i][0] = 0
			}
			dp[i][1] = dp[i-1][0] + 1
		}
		res = max(res, dp[i][0])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
