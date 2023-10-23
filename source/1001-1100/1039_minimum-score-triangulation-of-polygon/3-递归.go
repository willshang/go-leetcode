package main

import "math"

func main() {

}

var dp [][]int

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp = make([][]int, n) // dp[i][j]表示从i到j序列的最低分
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return dfs(values, 0, n-1)
}

func dfs(values []int, i, j int) int {
	if i+1 == j {
		return 0
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	res := math.MaxInt32
	for k := i + 1; k < j; k++ {
		sum := dfs(values, i, k) + dfs(values, k, j) + values[i]*values[k]*values[j]
		res = min(res, sum)
	}
	dp[i][j] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
