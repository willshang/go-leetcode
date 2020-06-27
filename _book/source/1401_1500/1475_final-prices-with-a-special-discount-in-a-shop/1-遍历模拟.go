package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}

// leetcode1475_商品折扣后的最终价格
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] = prices[i] - prices[j]
				break
			}
		}
	}
	return prices
}
