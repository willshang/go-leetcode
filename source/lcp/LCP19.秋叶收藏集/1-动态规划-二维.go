package main

import "fmt"

func main() {
	fmt.Println(minimumOperations("rrryyyrryyyrr"))
}

func minimumOperations(leaves string) int {
	n := len(leaves)
	// 长度i+1
	// dp[i][0] 全部变成r的步数
	// dp[i][1] 变成r...ry...y的步数
	// dp[i][2] 变成r...ry...yr...r的步数
	dp := make([][3]int, n)
	if leaves[0] == 'y' {
		dp[0][0] = 1 // 1个y变为r需要1步
	}
	for i := 1; i < n; i++ {
		if leaves[i] == 'r' {
			dp[i][0] = dp[i-1][0]     // 不需要改变，同前一个
			dp[i][1] = dp[i-1][0] + 1 // 全r+当前r，需要改变一个y，步数+1
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
				dp[i][2] = dp[i-1][1]
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2])
			}
		} else {
			dp[i][0] = dp[i-1][0] + 1 // 需要改变，步数+1
			dp[i][1] = dp[i-1][0]     // 前一个全r+当前y,不需要改变
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1])
				dp[i][2] = dp[i-1][1] + 1
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2]+1)
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
