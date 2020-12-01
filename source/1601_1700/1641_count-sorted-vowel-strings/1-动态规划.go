package main

func main() {

}

func countVowelStrings(n int) int {
	dp := make([][5]int, n+1)
	dp[0][0], dp[0][1], dp[0][2], dp[0][3], dp[0][4] = 1, 1, 1, 1, 1
	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] = dp[i][j] + dp[i-1][k]
			}
		}
	}
	res := 0
	for i := 0; i < 5; i++ {
		res = res + dp[n-1][i]
	}
	return res
}
