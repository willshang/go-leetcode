package main

import "math"

func main() {

}

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1) // dp[i][j] 左i右j的值
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, m+1)
	}
	res := math.MinInt32
	for k := 1; k <= m; k++ {
		for i := 0; i <= k; i++ {
			left := i
			right := k - i
			if i == 0 {
				dp[left][right] = dp[left][right-1] + nums[n-right]*multipliers[k-1]
			} else if i == k {
				dp[left][right] = dp[left-1][right] + nums[left-1]*multipliers[k-1]
			} else {
				l := dp[left][right-1] + nums[n-right]*multipliers[k-1]
				r := dp[left-1][right] + nums[left-1]*multipliers[k-1]
				dp[left][right] = max(l, r)
			}
			if k == m {
				res = max(res, dp[left][right])
			}
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
