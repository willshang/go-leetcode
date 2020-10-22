package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

var res int

func coinChange(coins []int, amount int) int {
	for i := 0; i < len(coins); i++ {
		for j := 0; j < len(coins)-1-i; j++ {
			if coins[j] < coins[j+1] {
				coins[j], coins[j+1] = coins[j+1], coins[j]
			}
		}
	}
	res = math.MaxInt32
	dfs(coins, amount, 0, 0)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func dfs(coins []int, amount int, count int, level int) {
	if amount == 0 {
		res = min(res, count)
		return
	}
	if level == len(coins) {
		return
	}
	for i := amount / coins[level]; i >= 0 && i+count < res; i-- {
		dfs(coins, amount-i*coins[level], count+i, level+1)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
