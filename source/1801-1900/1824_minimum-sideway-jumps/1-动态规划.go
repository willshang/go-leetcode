package main

import "math"

func main() {

}

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][3]int, n) // dp[i][j] 到达第i，下标为j的跑道最少次数
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = math.MaxInt32 / 10
		}
	}
	dp[0][0] = 1
	dp[0][1] = 0
	dp[0][2] = 1
	for i := 1; i < n; i++ {
		// 当前位置无障碍物，继承之前的次数，次数不变
		if obstacles[i] != 1 {
			dp[i][0] = dp[i-1][0]
		}
		if obstacles[i] != 2 {
			dp[i][1] = dp[i-1][1]
		}
		if obstacles[i] != 3 {
			dp[i][2] = dp[i-1][2]
		}
		// 从其它位置跳过来
		if obstacles[i] != 1 {
			dp[i][0] = min(dp[i][0], min(dp[i][1], dp[i][2])+1)
		}
		if obstacles[i] != 2 {
			dp[i][1] = min(dp[i][1], min(dp[i][0], dp[i][2])+1)
		}
		if obstacles[i] != 3 {
			dp[i][2] = min(dp[i][2], min(dp[i][0], dp[i][1])+1)
		}
	}
	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
