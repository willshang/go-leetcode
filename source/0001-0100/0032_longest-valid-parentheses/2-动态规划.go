package main

import "fmt"

func main() {
	//fmt.Println(longestValidParentheses("((())"))
	fmt.Println(longestValidParentheses("()((()))"))
}

func longestValidParentheses(s string) int {
	res := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			// '()' 匹配到
			if s[i-1] == '(' {
				if i < 2 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else {
				// '))'情况
				if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] < 2 {
						dp[i] = dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					}
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
