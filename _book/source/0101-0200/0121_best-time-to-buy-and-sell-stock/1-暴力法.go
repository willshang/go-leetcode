package main

import "fmt"

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
	fmt.Println(maxProfit([]int{1, 2}))
}

func maxProfit(prices []int) int {
	max := 0
	length := len(prices)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j <= length-1; j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
