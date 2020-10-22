package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

// leetcode44_通配符匹配
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		if p[i-1] == '*' { // 可以匹配任意字符串（包括空字符串）
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				// dp[i][j-1]=>不使用这个*，dp[i-1][j]=>使用这个*
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}
