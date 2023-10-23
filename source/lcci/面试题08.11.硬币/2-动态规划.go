package main

import "fmt"

func main() {
	fmt.Println(waysToChange(10))
}

// 程序员面试金典08.11_硬币
func waysToChange(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[n] % 1000000007
}
