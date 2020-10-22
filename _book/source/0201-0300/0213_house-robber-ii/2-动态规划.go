package main

import "fmt"

func main() {
	fmt.Println(rob([]int{0, 0}))
}

// leetcode213_打家劫舍II
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(getMax(nums[:n-1]), getMax(nums[1:]))
}

func getMax(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
