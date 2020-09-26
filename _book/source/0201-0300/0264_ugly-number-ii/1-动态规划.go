package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

// leetcode264_丑数II
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	// 丑数*2或3或5之后还是丑数
	idx2, idx3, idx5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(dp[idx2]*2, min(dp[idx3]*3, dp[idx5]*5))
		if dp[i] == dp[idx2]*2 {
			idx2++
		}
		if dp[i] == dp[idx3]*3 {
			idx3++
		}
		if dp[i] == dp[idx5]*5 {
			idx5++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
