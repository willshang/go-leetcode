package main

import "math"

func main() {

}

// leetcode1981_最小化目标值与所选元素的差
func minimizeTheDifference(mat [][]int, target int) int {
	n := len(mat)
	m := len(mat[0])
	maxValue := 5000
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, maxValue)
	}
	for j := 0; j < m; j++ {
		dp[0][mat[0][j]] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := mat[i][j]; k < maxValue; k++ {
				dp[i][k] = dp[i][k] | dp[i-1][k-mat[i][j]]
			}
		}
	}
	res := math.MaxInt32
	for j := 0; j < maxValue; j++ {
		if dp[n-1][j] == 1 {
			res = min(res, abs(j-target))
		}
	}
	return res
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
