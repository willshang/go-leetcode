package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Println(maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}

func maxCoins(piles []int) int {
	res := 0
	sort.Ints(piles)
	left := 0
	right := len(piles) - 1
	for left < right {
		left++
		right--
		res = res + piles[right]
		right--
	}
	return res
}
