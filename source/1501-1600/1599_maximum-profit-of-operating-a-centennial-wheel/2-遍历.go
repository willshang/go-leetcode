package main

import "fmt"

func main() {
	fmt.Println(minOperationsMaxProfit([]int{8, 3}, 5, 6))
}

// leetcode1599_经营摩天轮的最大利润
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	maxValue := 0
	res := -1
	total := 0
	i := 0
	profit := 0
	for total > 0 || i < len(customers) {
		if i < len(customers) {
			total = total + customers[i]
		}
		count := min(total, 4)
		total = total - count
		profit = profit + count*boardingCost - runningCost
		if profit > maxValue {
			maxValue = profit
			res = i + 1
		}
		i++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
