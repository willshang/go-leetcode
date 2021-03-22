package main

func main() {

}

// leetcode1771_由子序列构造的最长回文串的长度
func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	a := len(word1)
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	res := 0
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 // 内层+2
				if i < a && j >= a {
					res = max(res, dp[i][j])
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
