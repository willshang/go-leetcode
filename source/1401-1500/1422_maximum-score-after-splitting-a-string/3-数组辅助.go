package main

import "fmt"

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00"))
}

func maxScore(s string) int {
	max := 0
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			dp[i+1] = dp[i] + 1
		} else {
			dp[i+1] = dp[i]
		}
	}
	for i := 1; i < len(s); i++ {
		zero := i - dp[i]
		one := dp[len(s)] - dp[i]
		v := zero + one
		if v > max {
			max = v
		}
	}
	return max
}
