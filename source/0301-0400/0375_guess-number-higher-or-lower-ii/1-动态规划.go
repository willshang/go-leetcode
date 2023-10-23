package main

import "math"

func main() {

}

// leetcode375_猜数字大小II
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1) // 表示从[i,j]之间选择一个数字来猜，你确保获胜所需要的最少现金
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := n; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minValue := math.MaxInt32
			for k := i; k < j; k++ { // 可以选择[i-j]，其中最小值
				minValue = min(minValue, max(dp[i][k-1], dp[k+1][j])+k)
			}
			dp[i][j] = minValue
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
