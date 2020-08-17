package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(arr))
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	dp := make([]int, m)
	dp[0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < n; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < m; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[m-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
