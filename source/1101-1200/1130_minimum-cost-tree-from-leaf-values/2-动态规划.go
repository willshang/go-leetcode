package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))
}

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	maxArr := make([][]int, n)
	for i := 0; i < n; i++ {
		maxArr[i] = make([]int, n)
		maxValue := arr[i]
		for j := i; j < n; j++ {
			maxValue = max(maxValue, arr[j])
			maxArr[i][j] = maxValue // i到j之间的最大值
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			dp[i][j] = math.MaxInt32
			for k := i; k+1 <= j; k++ {
				// dp[i][j]代表从i到j之间的最小代价
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+maxArr[i][k]*maxArr[k+1][j])
			}
		}
	}
	return dp[0][n-1]
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
