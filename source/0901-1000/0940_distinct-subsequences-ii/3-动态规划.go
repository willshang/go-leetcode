package main

import "fmt"

func main() {
	fmt.Println(distinctSubseqII("abc"))
	fmt.Println(distinctSubseqII("aaa"))
	fmt.Println(distinctSubseqII("aba"))
}

var mod = 1000000007

func distinctSubseqII(s string) int {
	n := len(s)
	dp := make([]int, n+1) // dp[i]表示：长度为i的组合数
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if s[i-1] != s[j-1] {
				dp[i] = (dp[i] + dp[j]) % mod
			}
		}
		dp[i]++ // 当前dp[i]时：最长为i只有1种
	}
	res := 0
	for i := 1; i <= n; i++ {
		res = (res + dp[i]) % mod
	}
	return res
}
