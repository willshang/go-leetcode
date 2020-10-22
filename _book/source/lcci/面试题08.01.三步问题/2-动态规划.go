package main

import "fmt"

func main() {
	fmt.Println(waysToStep(3))
}

// 程序员面试金典08.01_三步问题
func waysToStep(n int) int {
	dp := make([]int, n+3)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	for i := 3; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}
	return dp[n-1]
}
