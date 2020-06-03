package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(arr))
}

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			left := 0
			up := 0
			if i > 0 {
				left = dp[j]
			}
			if j > 0 {
				up = dp[j-1]
			}
			// dp[j]=grid[i-1][j-1]+max(dp[j-1],dp[j])
			dp[j] = max(left, up) + grid[i][j]
		}
	}
	return dp[len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
