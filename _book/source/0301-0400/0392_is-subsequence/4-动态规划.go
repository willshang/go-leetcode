package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

/*
状态定义: dp[i][j] 表示长度为i的字符串s是否为长度为j的字符串t的子序列
状态转移方程: 如果s[i] == t[j]， 则dp[i][j] = dp[i-1][j-1]
如果s[i] != t[j]，则dp[i][j] = dp[i][j-1]
初始: dp[0][j] = true 表示空串s 是任意长度串t的子串
dp[i][0] = false 表示任意长度非空串s 不是空串t的字串
dp[i][0] = false 表示任意长度非空串s 不是空串t的字串
*/
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) > len(t) {
		return false
	}

	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(t)+1)
		dp[i][0] = false
	}
	for i := 0; i <= len(t); i++ {
		dp[0][i] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)]
}
