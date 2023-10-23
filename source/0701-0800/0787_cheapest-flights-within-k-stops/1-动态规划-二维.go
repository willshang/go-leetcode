package main

import "math"

func main() {

}

// leetcode787_K站中转内最便宜的航班
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	maxValue := math.MaxInt32 / 10
	dp := make([][]int, k+2) // dp[i][j] => 经过i次航班到j地需要的最少花费（k次中转需要k+1次航班）
	for i := 0; i <= k+1; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = maxValue
		}
	}
	dp[0][src] = 0 // 到开始地为0
	for i := 1; i <= k+1; i++ {
		for j := 0; j < len(flights); j++ {
			a, b, c := flights[j][0], flights[j][1], flights[j][2] // a=>b c
			dp[i][b] = min(dp[i][b], dp[i-1][a]+c)
		}
	}
	res := maxValue
	for i := 1; i <= k+1; i++ {
		res = min(res, dp[i][dst])
	}
	if res == maxValue {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
