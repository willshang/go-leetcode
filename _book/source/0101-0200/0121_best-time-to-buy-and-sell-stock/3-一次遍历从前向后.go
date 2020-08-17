package main

import "fmt"

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
	fmt.Println(maxProfit([]int{1, 2}))
}

// 从前向后
// leetcode 121 买卖股票的最佳时机
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
