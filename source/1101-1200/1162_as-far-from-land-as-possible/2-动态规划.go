package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	fmt.Println(maxDistance(arr))
}

func maxDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt32
			}
		}
	}
	// 从上往下
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				continue
			}
			if i >= 1 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			}
			if j >= 1 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	// 从下往上
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			if i < n-1 {
				dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
			}
			if j < m-1 {
				dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
			}
		}
	}
	res := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				continue
			}
			res = max(res, dp[i][j])
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
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
