package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

var dp [][]int

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	return dfs(s, p, 0, 0)
}

func dfs(s, p string, i, j int) bool {
	if i == len(s) && j == len(p) {
		return true
	}
	if dp[i][j] > 0 {
		if dp[i][j] == 1 {
			return false
		} else {
			return true
		}
	}
	if i >= len(s) {
		return p[j] == '*' && dfs(s, p, i, j+1)
	}
	if j >= len(p) {
		return false
	}
	res := false
	if p[j] == '*' {
		res = dfs(s, p, i+1, j) || dfs(s, p, i, j+1)
	} else {
		res = (s[i] == p[j] || p[j] == '?') && dfs(s, p, i+1, j+1)
	}
	if res == true {
		dp[i][j] = 2
	} else {
		dp[i][j] = 1
	}
	return res
}
