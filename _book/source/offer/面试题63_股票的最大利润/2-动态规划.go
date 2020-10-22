package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// 剑指offer_面试题63.股票的最大利润
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if profit < prices[i]-min {
			profit = prices[i] - min
		}
	}
	return profit
}
