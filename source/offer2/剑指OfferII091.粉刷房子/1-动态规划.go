package main

import "fmt"

func main() {
	fmt.Println(minCost([][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}))
}

// 剑指OfferII091.粉刷房子
func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][3]int, n) // dp[i][j] 表示涂前i间房子的最小成本
	for j := 0; j < 3; j++ {
		dp[0][j] = costs[0][j]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
