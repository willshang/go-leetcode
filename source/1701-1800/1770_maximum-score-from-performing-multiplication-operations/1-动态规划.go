package main

import "math"

func main() {

}

// leetcode1770_执行乘法运算的最大分数
func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1) // dp[i][j] 左i右j的值
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + nums[i-1]*multipliers[i-1] // 左i右0
		dp[0][i] = dp[0][i-1] + nums[n-i]*multipliers[i-1] // 左0右i
	}
	for i := 1; i <= m; i++ {
		for j := 1; i+j <= m; j++ {
			left := dp[i-1][j] + nums[i-1]*multipliers[i+j-1]
			right := dp[i][j-1] + nums[n-j]*multipliers[i+j-1]
			dp[i][j] = max(left, right)
		}
	}
	res := math.MinInt32
	for i := 0; i <= m; i++ {
		res = max(res, dp[i][m-i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
