package main

import "fmt"

func main() {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(arr))
}

var dp [][]int

func minimumTotal(triangle [][]int) int {
	dp = make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle))
	}
	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	dp[i][j] = min(dfs(triangle, i+1, j), dfs(triangle, i+1, j+1)) + triangle[i][j]
	return dp[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
