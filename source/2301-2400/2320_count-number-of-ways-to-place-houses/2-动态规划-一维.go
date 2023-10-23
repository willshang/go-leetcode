package main

func main() {

}

var mod = 1000000007

func countHousePlacements(n int) int {
	dp := make([]int, n+3) // 只考虑单侧的方案
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= n; i++ {
		// 不放；dp[i] = dp[i-1]
		// 放: i-1不能放，dp[i] = dp[i-2]
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return dp[n] * dp[n] % mod // 考虑2侧情况
}
