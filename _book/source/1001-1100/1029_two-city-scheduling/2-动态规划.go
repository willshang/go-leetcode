package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
}

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := i + 1; j <= n; j++ {
			// dp[i][j]表示i个人飞往A市的次数为j的最低费用
			// 无效掉j>i的情况，比如i=3, j=4
			// 因为不存在3个人飞往A市次数为4次的情况
			dp[i][j] = 100000000
		}
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + costs[i-1][1]
		for j := 1; j <= i; j++ {
			// dp[i][j]表示i个人飞往A市的次数为j的最低费用
			// 其中i-1个人飞往A市的次数为j+当前飞往B市的费用
			// 其中i-1个人飞往A市的次数为j-1+当前飞往A市的费用
			dp[i][j] = min(dp[i-1][j]+costs[i-1][1], dp[i-1][j-1]+costs[i-1][0])
		}
	}
	return dp[n][n/2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
