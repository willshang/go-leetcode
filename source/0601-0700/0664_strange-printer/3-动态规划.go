package main

func main() {

}

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n+1) // dp[i][j] => 打印S[i], S[i+1], ..., S[j]所需的次数
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][i] = 1
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			dp[i][j] = dp[i+1][j] + 1
			for k := i + 1; k <= j; k++ {
				if s[i] == s[k] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
