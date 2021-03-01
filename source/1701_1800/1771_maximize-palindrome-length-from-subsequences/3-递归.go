package main

func main() {

}

var dp [][]int
var a int
var res int

func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	a = len(word1)
	res = 0
	n := len(s)
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dfs(s, 0, n-1)
	return res
}

func dfs(s string, i, j int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	if s[i] == s[j] {
		dp[i][j] = dfs(s, i+1, j-1) + 2
		if i < a && j >= a {
			res = max(res, dp[i][j])
		}
	} else {
		dp[i][j] = max(dfs(s, i+1, j), dfs(s, i, j-1))
	}
	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
