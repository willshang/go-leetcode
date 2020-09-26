package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

var dp [][]int

func longestPalindromeSubseq(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	n := len(s)
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return dfs(s, 0, n-1)
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
