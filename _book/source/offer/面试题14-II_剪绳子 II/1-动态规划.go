package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
}

func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+10)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	dp[4] = 4
	dp[5] = 6
	dp[6] = 9
	dp[7] = 12
	for i := 8; i <= n; i++ {
		dp[i] = (dp[i-3] * 3) % 1000000007
	}
	return dp[n]
}
