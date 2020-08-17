package main

import "fmt"

func main() {
	var arr interface{}
	arr = []string{"a", "b"}
	fmt.Println(arr)
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

// leetcode309_最佳买卖股票时机含冷冻期
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][3]int, n)
	// 0 => 持有
	// 1 => 不持有，本日卖出，下一日冷冻期
	// 2 => 不持有，本日无卖出，下一日不是冷冻期
	dp[0][0] = -prices[0] // 第0天买入，亏损-price[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) // 继续持有 or 可以操作，继续买入，导致今天持有股票
		dp[i][1] = dp[i-1][0] + prices[i]                // 卖出操作
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])           // 昨天卖出，无股票，今天是冷冻期 or 昨天没股票，也不操作
	}
	return max(dp[n-1][1], dp[n-1][2]) // 最后一天操作，会导致利润变少，可以忽略
	// return max(dp[n-1][0], max(dp[n-1][1], dp[n-1][2]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
