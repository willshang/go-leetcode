package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

// dp(l,r)=dp(l+1,r−1)&&(s[l]==s[r])
// dp[l,r]：字符串s从索引l到r的子串是否是回文串
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	start := 0
	max := 1
	for r := 0; r < len(s); r++ {
		dp[r] = make([]bool, len(s))
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
			if dp[l][r] == true {
				if r-l+1 > max {
					max = r - l + 1
					start = l
				}
			}
		}
	}
	return s[start : start+max]
}
