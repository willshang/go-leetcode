package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minimumOperations("rrryyyrryyyrr"))
	fmt.Println(minimumOperations("yry"))
}

// leetcode_lcp19_秋叶收藏集
func minimumOperations(leaves string) int {
	n := len(leaves)
	// 长度i+1
	// dp[i][0] 全部变成r的步数
	// dp[i][1] 变成r...ry...y的步数
	// dp[i][2] 变成r...ry...yr...r的步数
	dp := make([][3]int, n)
	maxValue := math.MaxInt32 / 10
	if leaves[0] == 'y' {
		dp[0][0] = 1 // 1个y变为r需要1步
		dp[0][1] = maxValue
		dp[0][2] = maxValue
	}
	for i := 1; i < n; i++ {
		dp[i][2] = maxValue
		dp[i][1] = maxValue
		if leaves[i] == 'r' {
			dp[i][0] = dp[i-1][0]                      // 不需要改变，同前一个
			dp[i][1] = min(dp[i-1][0]+1, dp[i-1][1]+1) // 前一个r+1，前一个y+1
			if i >= 2 {
				dp[i][2] = min(dp[i-1][1], dp[i-1][2]) // 前一个y不变，前一个r不变
			}
		} else {
			dp[i][0] = dp[i-1][0] + 1              // 需要改变，步数+1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) // 前一个r不变，前一个y不变
			if i >= 2 {
				dp[i][2] = min(dp[i-1][1]+1, dp[i-1][2]+1) // 前一个y+1, 前一个r+1
			}
		}
	}
	return dp[n-1][2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
