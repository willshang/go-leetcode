package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

/*
dp 公式, 2个状态，dp[i][0] 表示第i天不抢，dp[i][1]表示第i天抢
dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
dp[i][1] = nums[i] + dp[i-1][0]
*/
// 保存二维dp数组
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp := make([][]int, n)
	for n := range dp {
		dp[n] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
