package main

func main() {

}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	// dp[i][j]经过第i个加油站加油j次能够到达的最远距离
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = startFuel
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 不加油
			if dp[i-1][j] >= stations[i-1][0] {
				dp[i][j] = dp[i-1][j]
			}
			// 加油
			if dp[i-1][j-1] >= stations[i-1][0] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+stations[i-1][1])
			}
		}
	}
	for i := 0; i <= n; i++ {
		if dp[n][i] >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
