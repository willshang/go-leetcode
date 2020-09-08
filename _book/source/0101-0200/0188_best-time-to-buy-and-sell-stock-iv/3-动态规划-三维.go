package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
}

// leetcode188_买卖股票的最佳时机IV
func maxProfit(k int, prices []int) int {
	res := 0
	if k >= len(prices)/2 {
		for i := 0; i < len(prices)-1; i++ {
			if prices[i] < prices[i+1] {
				res = res + prices[i+1] - prices[i]
			}
		}
		return res
	}
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 0; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
