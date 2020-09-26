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
	count := len(piles) / 3
	for i := 0; i < count; i++ {
		index := len(piles) - 1 - i*2 - 1
		res = res + piles[index]
	}
	return res
}
