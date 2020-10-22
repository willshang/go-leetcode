package main

import "fmt"

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			length := dp[j] * dp[i-j]
			if length > max {
				max = length
			}
			dp[i] = max
		}
	}
	return dp[n]
}
