package main

import "fmt"

func main() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
				continue
			}
			minValue := dp[i-1][j]
			if j > 0 {
				minValue = min(minValue, dp[i-1][j-1])
			}
			if j < n-1 {
				minValue = min(minValue, dp[i-1][j+1])
			}
			dp[i][j] = dp[i][j] + minValue + matrix[i][j]
		}
	}
	res := dp[n-1][0]
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
