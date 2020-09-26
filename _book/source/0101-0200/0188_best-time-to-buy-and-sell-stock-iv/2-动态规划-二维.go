package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
}

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
	dp0, dp1 := make([][]int, len(prices)), make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp0[i] = make([]int, k+1)
		dp1[i] = make([]int, k+1)
	}
	for i := 0; i <= k; i++ {
		dp1[0][i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp0[i][j] = max(dp0[i-1][j], dp1[i-1][j]+prices[i])
			dp1[i][j] = max(dp1[i-1][j], dp0[i-1][j-1]-prices[i])
		}
	}
	return dp0[len(prices)-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
