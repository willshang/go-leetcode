package main

func main() {

}

// leetcode2008_出租车的最大盈利
func maxTaxiEarnings(n int, rides [][]int) int64 {
	dp := make([]int, n+1) // dp[i]到达i位置的最大盈利
	arr := make([][][2]int, n+1)
	for i := 0; i < len(rides); i++ {
		a, b, c := rides[i][0], rides[i][1], rides[i][2]
		arr[b] = append(arr[b], [2]int{a, b - a + c})
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		for j := 0; j < len(arr[i]); j++ {
			a, c := arr[i][j][0], arr[i][j][1]
			dp[i] = max(dp[i], dp[a]+c)
		}
	}
	return int64(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
