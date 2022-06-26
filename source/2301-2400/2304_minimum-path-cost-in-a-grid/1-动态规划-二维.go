package main

import "math"

func main() {

}

// leetcode2304_网格中的最小路径代价
func minPathCost(grid [][]int, moveCost [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		dp[0][j] = grid[0][j] // 第一行的代价
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			value := math.MaxInt32
			for k := 0; k < m; k++ {
				// 上一行的代价+当前的单元值+移动的代价
				prev := grid[i-1][k] // 上一层的值
				value = min(value, dp[i-1][k]+grid[i][j]+moveCost[prev][j])
			}
			dp[i][j] = value
		}
	}
	res := math.MaxInt32
	for j := 0; j < m; j++ {
		res = min(res, dp[n-1][j])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
