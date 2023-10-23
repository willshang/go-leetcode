package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMoneyAmount(10))
}

var dp [][]int

func getMoneyAmount(n int) int {
	dp = make([][]int, n+1) // 表示从[i,j]之间选择一个数字来猜，你确保获胜所需要的最少现金
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	return dfs(1, n)
}

func dfs(start, end int) int {
	if start >= end {
		return 0
	}
	if dp[start][end] > 0 {
		return dp[start][end]
	}
	minValue := math.MaxInt32
	for i := start; i <= end; i++ {
		res := i + max(dfs(i+1, end), dfs(start, i-1))
		minValue = min(minValue, res)
	}
	dp[start][end] = minValue
	return minValue
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
