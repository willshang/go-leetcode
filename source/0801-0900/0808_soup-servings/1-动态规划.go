package main

func main() {

}

// leetcode808_分汤
func soupServings(N int) float64 {
	n := N / 25
	if N%25 > 0 {
		n = n + 1
	}
	if n >= 500 {
		return 1.0
	}
	// 当给定i毫升的A和j毫升的B的概率
	// dp[i][j]的概率=0.25*(dp[i-4][j]+dp[i-3][j-1]+dp[i-2][j-2]+dp[i-1][j-3])
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
	}
	dp[0][0] = 0.5
	for i := 1; i <= n; i++ {
		dp[i][0] = 0
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		a1 := max(i-4, 0)
		a2 := max(i-3, 0)
		a3 := max(i-2, 0)
		a4 := max(i-1, 0)
		for j := 1; j <= n; j++ {
			b1 := max(j, 0)
			b2 := max(j-1, 0)
			b3 := max(j-2, 0)
			b4 := max(j-3, 0)
			dp[i][j] = 0.25 * (dp[a1][b1] + dp[a2][b2] + dp[a3][b3] + dp[a4][b4])
		}
	}
	return dp[n][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
