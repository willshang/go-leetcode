package main

import "math"

func main() {

}

// leetcode1824_最少侧跳次数
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
		if obstacles[i] == 0 { // 没有障碍物，从其它2条道跳过来
			dp[i][0] = min(dp[i-1][0], min(dp[i-1][1], dp[i-1][2])+1)
			dp[i][1] = min(dp[i-1][1], min(dp[i-1][0], dp[i-1][2])+1)
			dp[i][2] = min(dp[i-1][2], min(dp[i-1][0], dp[i-1][1])+1)
		} else {
			a := obstacles[i] - 1
			b := (obstacles[i]) % 3
			c := (obstacles[i] + 1) % 3
			dp[i][a] = math.MaxInt32 / 10 // 不可达
			dp[i][b] = min(dp[i-1][b], dp[i-1][c]+1)
			dp[i][c] = min(dp[i-1][c], dp[i-1][b]+1)
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
