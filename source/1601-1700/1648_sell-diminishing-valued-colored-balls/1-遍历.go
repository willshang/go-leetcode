package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProfit([]int{3, 7}, 10))
}

// leetcode1648_销售价值减少的颜色球
func maxProfit(inventory []int, orders int) int {
	inventory = append(inventory, 0) // 避免第一个数特判
	sort.Ints(inventory)
	n := len(inventory)
	res := 0
	// 每次把当前数减到前一个数
	for i := n - 1; i >= 1; i-- {
		if orders <= 0 {
			break
		}
		total := (n - i) * (inventory[i] - inventory[i-1])
		if total <= orders { // 够用
			sum := (inventory[i-1] + 1 + inventory[i]) * (inventory[i] - inventory[i-1]) / 2 * (n - i)
			res = (res + sum) % 1000000007
			orders = orders - total
		} else { // 不够用
			a, b := orders/(n-i), orders%(n-i)
			start := inventory[i] - a + 1
			sum := (start+inventory[i])*a/2*(n-i) + b*(start-1)
			res = (res + sum) % 1000000007
			orders = 0
		}
	}
	return res
}
