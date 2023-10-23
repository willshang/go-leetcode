package main

func main() {

}

// leetcode2320_统计放置房子的方式数
var mod = 1000000007

func countHousePlacements(n int) int {
	dp := make([][4]int, n+1)
	dp[1][0] = 1 // 2侧没有房子
	dp[1][1] = 1 // 1侧有房子：上
	dp[1][2] = 1 // 1侧有房子：下
	dp[1][3] = 1 // 2侧都有房子
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3]) % mod
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][3] = (dp[i-1][0]) % mod
	}
	sum := 0
	for i := 0; i < 4; i++ {
		sum = (sum + dp[n][i]) % mod
	}
	return sum
}
