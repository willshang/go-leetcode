package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(arr))
}

var dp [][]int

func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	return dfs(dungeon, n, m, 0, 0)
}

func dfs(dungeon [][]int, n, m, i, j int) int {
	if i == n-1 && j == m-1 {
		return max(1-dungeon[i][j], 1)
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	res := 0
	if i == n-1 {
		res = max(dfs(dungeon, n, m, i, j+1)-dungeon[i][j], 1)
	} else if j == m-1 {
		res = max(dfs(dungeon, n, m, i+1, j)-dungeon[i][j], 1)
	} else {
		minValue := min(dfs(dungeon, n, m, i, j+1), dfs(dungeon, n, m, i+1, j))
		res = max(minValue-dungeon[i][j], 1)
	}
	dp[i][j] = res
	return dp[i][j]
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
