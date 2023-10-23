package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

func isMatch(s string, p string) bool {
	// dp[i][j]表示p[:i]能否正则匹配s[:j]
	dp := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	// 1.初始化
	dp[0][0] = true
	for i := 2; i < len(p)+1; i++ {
		if i%2 == 0 && p[i-1] == '*' {
			dp[i][0] = dp[i-2][0]
		}
	}
	// 2.dp状态转移
	for i := 1; i < len(p)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			// 2.1 相同或者 .
			if p[i-1] == s[j-1] || p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == '*' {
				if i > 1 {
					if p[i-2] == s[j-1] || p[i-2] == '.' {
						dp[i][j] = dp[i][j-1] || dp[i-2][j-1] || dp[i-2][j]
					} else {
						dp[i][j] = dp[i-2][j]
					}
				}
			}
		}
	}
	return dp[len(p)][len(s)]
}
