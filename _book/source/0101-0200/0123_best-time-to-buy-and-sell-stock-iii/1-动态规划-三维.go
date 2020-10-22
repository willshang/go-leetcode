package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}

// leetcode123_买卖股票的最佳时机III
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxK := 2
	// 第一维n个状态：n天
	// 第二维3个状态：0、1、2分别表示完成的交易次数
	// 第三维2个状态：0（不持有股票）、1（持有股票）
	dp := make([][3][2]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= maxK; j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j][1] = math.MinInt64
				} else {
					dp[i][j][1] = -prices[i]
				}
			} else if j == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = math.MinInt64
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][2][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
