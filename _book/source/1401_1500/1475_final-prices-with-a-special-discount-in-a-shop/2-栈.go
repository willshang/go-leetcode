package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}

func finalPrices(prices []int) []int {
	stack := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		for len(stack) > 0 {
			index := stack[len(stack)-1]
			if prices[i] > prices[index] {
				break
			}
			prices[index] = prices[index] - prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}
