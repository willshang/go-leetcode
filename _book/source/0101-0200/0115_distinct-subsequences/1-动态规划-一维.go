package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		for j := len(t); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[len(t)]
}
