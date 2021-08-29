package main

import "math"

func main() {

}

// leetcode1928_规定时间内到达终点的最小花费
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	maxValue := math.MaxInt32 / 10
	n := len(passingFees)
	dp := make([][]int, maxTime+1) // dp[i][j] => 在i分钟到达j城市的最少花费
	for i := 0; i <= maxTime; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = maxValue
		}
	}
	dp[0][0] = passingFees[0] // 出发的城市也要收费
	for i := 1; i <= maxTime; i++ {
		for j := 0; j < len(edges); j++ {
			a, b, c := edges[j][0], edges[j][1], edges[j][2] // a=>b c
			if c <= i {                                      // 小于时间i
				// 注意：无向图
				dp[i][a] = min(dp[i][a], dp[i-c][b]+passingFees[a])
				dp[i][b] = min(dp[i][b], dp[i-c][a]+passingFees[b])
			}
		}
	}
	res := maxValue
	for i := 1; i <= maxTime; i++ {
		res = min(res, dp[i][n-1])
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
