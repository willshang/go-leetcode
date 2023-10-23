package main

func main() {

}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i <= len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= len(strs); i++ {
		a, b := getCount(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if a <= j && b <= k {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-a][k-b]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

func getCount(str string) (a, b int) {
	a, b = 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			a++
		} else {
			b++
		}
	}
	return a, b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
