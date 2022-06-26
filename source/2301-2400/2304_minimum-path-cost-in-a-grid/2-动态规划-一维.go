package main

import "math"

func main() {

}

func minPathCost(grid [][]int, moveCost [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([]int, m)
	copy(dp, grid[0])
	for i := 1; i < n; i++ {
		temp := make([]int, m)
		for j := 0; j < m; j++ {
			value := math.MaxInt32
			for k := 0; k < m; k++ {
				// 上一行的代价+当前的单元值+移动的代价
				prev := grid[i-1][k] // 上一层的值
				value = min(value, dp[k]+grid[i][j]+moveCost[prev][j])
			}
			temp[j] = value
		}
		copy(dp, temp)
	}
	res := math.MaxInt32
	for j := 0; j < m; j++ {
		res = min(res, dp[j])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
