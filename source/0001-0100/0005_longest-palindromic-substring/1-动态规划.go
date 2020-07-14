package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

// dp(i,j)=dp(i+1,j−1)∧(s[i]==s[j])
func longestPalindrome(s string) string {
	res := ""
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for l := 0; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if i == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(res) {
				res = s[i : i+1+l]
			}
		}
	}
	return res
}
