package main

import "math"

func main() {

}

func minimizeTheDifference(mat [][]int, target int) int {
	n := len(mat)
	m := len(mat[0])
	maxSum := 0
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, 5000)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		maxValue := 0
		for j := 0; j < m; j++ {
			maxValue = max(maxValue, mat[i][j])
			// 枚举起点+终点
			for k := mat[i][j]; k <= maxSum+mat[i][j]; k++ {
				if dp[i][k-mat[i][j]] == true { // 可以转移
					dp[i+1][k] = true
				}
			}
		}
		maxSum = maxSum + maxValue
	}
	res := math.MaxInt32
	for j := 0; j <= maxSum; j++ {
		if dp[n][j] == true {
			res = min(res, abs(j-target))
		}
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
