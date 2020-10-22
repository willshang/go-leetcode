package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	buy1, buy2 := math.MaxInt32, math.MaxInt32
	profit1, profit2 := 0, 0
	for i := 0; i < len(prices); i++ {
		value := prices[i]
		buy1 = min(buy1, value)
		profit1 = max(profit1, value-buy1)
		buy2 = min(buy2, value-profit1)
		profit2 = max(profit2, value-buy2)
	}
	return profit2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
