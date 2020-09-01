package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

func longestPalindromeSubseq(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		prev := 0
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = prev + 2 // 内层+2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = temp
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
