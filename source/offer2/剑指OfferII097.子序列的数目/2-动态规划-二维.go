package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

// 剑指OfferII097.子序列的数目
func numDistinct(s string, t string) int {
	// dp[i][j]为使用s的前i个字符能够最多组成多少个t的前j个字符
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				// s用最后一位的 +不用最后一位
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
