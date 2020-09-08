package main

import (
	"fmt"
	"math"
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
	dp0, dp1 := make([]int, k+1), make([]int, k+1)
	for i := 0; i <= k; i++ {
		dp0[i] = 0
		dp1[i] = math.MinInt64
	}
	for i := 0; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			dp0[j] = max(dp0[j], dp1[j]+prices[i])
			dp1[j] = max(dp1[j], dp0[j-1]-prices[i])
		}
	}
	return dp0[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
