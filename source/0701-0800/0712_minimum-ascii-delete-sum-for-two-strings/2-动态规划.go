package main

func main() {

}

func minimumDeleteSum(s1 string, s2 string) int {
	a, b := len(s1), len(s2)
	dp := make([][]int, a+1)
	for i := 0; i <= a; i++ {
		dp[i] = make([]int, b+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for i := 1; i <= b; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+int(s2[j-1]), dp[i-1][j]+int(s1[i-1]))
			}
		}
	}
	return dp[a][b]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
