package main

import "fmt"

func main() {
	fmt.Println(getKthMagicNumber(10))
}

// 程序员面试金典17.09_第k个数
func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	// *3或5或7之后得到
	idx3, idx5, idx7 := 0, 0, 0
	for i := 1; i < k; i++ {
		dp[i] = min(dp[idx3]*3, min(dp[idx5]*5, dp[idx7]*7))
		if dp[i] == dp[idx3]*3 {
			idx3++
		}
		if dp[i] == dp[idx5]*5 {
			idx5++
		}
		if dp[i] == dp[idx7]*7 {
			idx7++
		}
	}
	return dp[k-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
