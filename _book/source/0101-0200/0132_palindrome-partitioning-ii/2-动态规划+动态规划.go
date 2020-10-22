package main

import "fmt"

func main() {
	fmt.Println(minCut("aab"))
}

func minCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = -1
	dp[1] = 1
	arr := getDP(s)
	for i := 1; i <= len(s); i++ {
		dp[i] = i - 1 // 长度N切分n-1次
		for j := 0; j < i; j++ {
			if arr[j][i-1] == true {
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

func getDP(s string) [][]bool {
	dp := make([][]bool, len(s))
	for r := 0; r < len(s); r++ {
		dp[r] = make([]bool, len(s))
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
		}
	}
	return dp
}
