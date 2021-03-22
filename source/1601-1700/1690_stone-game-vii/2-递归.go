package main

import "fmt"

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2}))
}

var dp [][]int
var arr []int

func stoneGameVII(stones []int) int {
	n := len(stones)
	arr = make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + stones[i]
	}
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dp[i][j] = -1
		}
	}
	dfs(0, n-1)
	return dp[0][n-1]
}

func dfs(i, j int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	left := arr[j+1] - arr[i+1] - dfs(i+1, j)
	right := arr[j] - arr[i] - dfs(i, j-1)
	dp[i][j] = max(left, right)
	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
