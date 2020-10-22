package main

import "fmt"

func main() {
	var arr interface{}
	arr = []string{"a", "b"}
	fmt.Println(arr)
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// a => 持有
	// b => 不持有，本日卖出，下一日冷冻期
	// c => 不持有，本日无卖出，下一日不是冷冻期
	var a, b, c int
	a = -prices[0] // 第0天买入，亏损-price[0]
	for i := 1; i < n; i++ {
		A := max(a, c-prices[i]) // 继续持有 or 可以操作，继续买入，导致今天持有股票
		B := a + prices[i]       // 卖出操作
		C := max(b, c)           // 昨天卖出，无股票，今天是冷冻期 or 昨天没股票，也不操作
		a, b, c = A, B, C
	}
	return max(b, c)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
