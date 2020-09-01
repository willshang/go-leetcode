package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

// leetcode516_最长回文子序列
func longestPalindromeSubseq(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 // 内层+2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
