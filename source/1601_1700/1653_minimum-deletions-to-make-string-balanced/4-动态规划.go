package main

func main() {

}

func minimumDeletions(s string) int {
	n := len(s)
	dp := make([][2]int, n+1)
	// dp[n][0]以a结尾需要删除的次数
	// dp[n][1]以b结尾需要删除的次数
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			dp[i+1][0] = dp[i][0]
			dp[i+1][1] = dp[i][1] + 1
		} else {
			dp[i+1][0] = dp[i][0] + 1
			dp[i+1][1] = min(dp[i][0], dp[i][1])
		}
	}
	return min(dp[n][0], dp[n][1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
