package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(updateMatrix(arr))
}

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = math.MaxInt32 / 10
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if dp[i][j] > 1 {
				if i < n-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
				if j < m-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}
	return dp
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
