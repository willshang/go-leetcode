package main

import "math"

func main() {

}

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
	arr := make([][][2]int, n)
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1], edges[i][2] // a=>b c
		arr[a] = append(arr[a], [2]int{b, c})
		arr[b] = append(arr[b], [2]int{a, c})
	}
	dp[0][0] = passingFees[0] // 出发的城市也要收费
	for i := 1; i <= maxTime; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < len(arr[j]); k++ {
				a := j
				b, c := arr[j][k][0], arr[j][k][1]
				if c <= i { // 小于时间i
					// 注意：无向图
					dp[i][a] = min(dp[i][a], dp[i-c][b]+passingFees[a])
					dp[i][b] = min(dp[i][b], dp[i-c][a]+passingFees[b])
				}
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
