package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Println(maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}

// leetcode1561_你可以获得的最大硬币数目
func maxCoins(piles []int) int {
	res := 0
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})
	for i := 0; i < len(piles)/3; i++ {
		res = res + piles[i*2+1]
	}
	return res
}
