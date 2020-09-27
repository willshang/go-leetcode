package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(arr))
}

// leetcode174_地下城游戏
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1 // 结果最小为1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			minValue := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minValue-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
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
