package main

import "fmt"

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))

	fmt.Println(maxProfit([]int{1, 2}))
}
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max_sell_price := 0
	max_profit := 0

	for i := len(prices) - 1; i >= 0; i-- {
		if max_sell_price < prices[i] {
			max_sell_price = prices[i]
		}

		if max_profit < max_sell_price-prices[i] {
			max_profit = max_sell_price - prices[i]
		}
	}

	return max_profit
}

/*func maxProfit(prices []int) int {
	max := 0
	temp := 0
	for i := 1;i < len(prices); i++{
		temp = temp + prices[i] - prices[i-1]

		if temp < 0{
			temp = 0
		}
		if max < temp{
			max = temp
		}
	}
	return max
}*/
/*func maxProfit(prices []int) int {
	max := 0
	length := len(prices)

	for i := 0; i < length-1 ; i++{
		for j := i+1; j <= length-1; j++{
			if prices[j] - prices[i] > max{
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}*/
