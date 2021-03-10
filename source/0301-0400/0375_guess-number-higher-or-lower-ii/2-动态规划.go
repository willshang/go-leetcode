package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMoneyAmount(10))
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1) // 表示从[i,j]之间选择一个数字来猜，你确保获胜所需要的最少现金
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 2; i <= n; i++ {
		for j := 1; j+i-1 <= n; j++ {
			minValue := math.MaxInt32
			for k := j; k < i+j-1; k++ { // 可以选择[i-j]，其中最小值
				minValue = min(minValue, max(dp[j][k-1], dp[k+1][i+j-1])+k)
			}
			dp[j][i+j-1] = minValue
		}
	}
	return dp[1][n]
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
