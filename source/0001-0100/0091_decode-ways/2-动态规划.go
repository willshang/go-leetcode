package main

import "fmt"

func main() {
	fmt.Println(numDecodings("226"))
}

// leetcode91_解码方法
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if i == 0 || s[i-1] == '1' || s[i-1] == '2' {
				return 0
			}
			dp[i+1] = dp[i-1]
		} else {
			if i > 0 && (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
				dp[i+1] = dp[i-1] + dp[i]
			} else {
				dp[i+1] = dp[i]
			}
		}
	}
	return dp[len(s)]
}
