package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}

// leetcode518_零钱兑换II
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1 // 金额为0的情况，只有都不选，组合情况为1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}
