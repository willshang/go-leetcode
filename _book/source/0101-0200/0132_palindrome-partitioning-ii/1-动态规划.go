package main

import "fmt"

func main() {
	fmt.Println(minCut("aab"))
}

// leetcode132_分割回文串II
func minCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = -1
	dp[1] = 1
	for i := 1; i <= len(s); i++ {
		dp[i] = i - 1 // 长度N切分n-1次
		for j := 0; j < i; j++ {
			if judge(s[j:i]) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func judge(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
