package main

import "math"

func main() {

}

func maxDotProduct(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			value := nums1[i-1] * nums2[j-1] // 单独选对应的i,j 乘积
			dp[i][j] = max(dp[i][j], value)  // 至少选一对
			dp[i][j] = max(dp[i][j], dp[i-1][j-1]+value)
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			dp[i][j] = max(dp[i][j], dp[i][j-1])
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
