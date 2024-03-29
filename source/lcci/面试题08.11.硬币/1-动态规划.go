package main

import "fmt"

func main() {
	fmt.Println(waysToChange(10))
}

func waysToChange(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([][]int, 5)
	for i := 0; i <= 4; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1 // 金额为0的情况，只有都不选，组合情况为1
	}
	for i := 1; i <= 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[4][n] % 1000000007
}
