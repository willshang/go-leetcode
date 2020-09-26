package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfit(prices []int, fee int) int {
	dp0, dp1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		temp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, temp-prices[i]-fee)
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
