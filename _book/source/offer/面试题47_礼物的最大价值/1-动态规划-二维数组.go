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

// 剑指offer_面试题47_礼物的最大价值
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			left := 0
			up := 0
			if i > 0 {
				left = dp[i-1][j]
			}
			if j > 0 {
				up = dp[i][j-1]
			}
			// dp[i][j]=grid[i-1][j-1]+max(dp[i][j-1],dp[i-1][j])
			dp[i][j] = max(left, up) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
