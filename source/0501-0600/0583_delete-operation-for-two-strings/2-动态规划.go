package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	a, b := len(word1), len(word2)
	dp := make([][]int, a+1)
	for i := 0; i <= a; i++ {
		dp[i] = make([]int, b+1)
		dp[i][0] = i
	}
	for i := 0; i <= b; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + 1
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
