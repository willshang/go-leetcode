package main

import "math"

func main() {

}

func minimizeTheDifference(mat [][]int, target int) int {
	n := len(mat)
	m := len(mat[0])
	dp := make([]bool, target)
	dp[0] = true
	maxValue := math.MaxInt32 // 存储大于等于target的数
	for i := 0; i < n; i++ {
		temp := make([]bool, target)
		tempMaxValue := math.MaxInt32
		for j := 0; j < m; j++ {
			// 枚举起点+终点
			for k := 0; k < target; k++ {
				if dp[k] == true {
					if k+mat[i][j] >= target {
						tempMaxValue = min(tempMaxValue, k+mat[i][j])
					} else {
						temp[k+mat[i][j]] = true
					}
				}
			}
			if maxValue != math.MaxInt32 {
				tempMaxValue = min(tempMaxValue, maxValue+mat[i][j])
			}
		}
		dp = temp
		maxValue = tempMaxValue
	}
	res := abs(maxValue - target)
	for i := target - 1; i >= 0; i-- {
		if dp[i] == true {
			res = min(res, target-i) // i都小于target
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
