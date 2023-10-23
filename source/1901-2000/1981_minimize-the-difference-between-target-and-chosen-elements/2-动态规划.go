package main

import "math"

func main() {

}

func minimizeTheDifference(mat [][]int, target int) int {
	n := len(mat)
	m := len(mat[0])
	maxSum := 0
	dp := []int{1}
	for i := 0; i < n; i++ {
		maxValue := 0
		for j := 0; j < m; j++ {
			maxValue = max(maxValue, mat[i][j])
		}
		temp := make([]int, maxSum+maxValue+1)
		for j := 0; j < m; j++ {
			// 枚举起点+终点
			for k := mat[i][j]; k <= maxSum+mat[i][j]; k++ {
				temp[k] = temp[k] | dp[k-mat[i][j]]
			}
		}
		dp = temp
		maxSum = maxSum + maxValue
	}
	res := math.MaxInt32
	for j := 0; j <= maxSum; j++ {
		if dp[j] > 0 {
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
