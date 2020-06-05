package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	profit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if max < prices[i] {
			max = prices[i]
		}
		if profit < max-prices[i] {
			profit = max - prices[i]
		}
	}
	return profit
}
