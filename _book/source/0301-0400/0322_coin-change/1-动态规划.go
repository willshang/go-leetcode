package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

// leetcode322_零钱兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for j := 0; j < len(coins); j++ {
			prev := i - coins[j]
			if i < coins[j] || dp[prev] == -1 {
				continue
			}
			if dp[i] == -1 || dp[i] > dp[prev]+1 {
				dp[i] = dp[prev] + 1
			}
		}
	}
	return dp[amount]
}
