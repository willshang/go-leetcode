package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

// 剑指OfferII073.狒狒吃香蕉
func minEatingSpeed(piles []int, h int) int {
	maxValue := piles[0]
	for i := 1; i < len(piles); i++ {
		maxValue = max(maxValue, piles[i])
	}
	return sort.Search(maxValue, func(speed int) bool {
		if speed == 0 { // 避免除0
			return false
		}
		total := 0
		for i := 0; i < len(piles); i++ {
			total = total + piles[i]/speed
			if piles[i]%speed > 0 {
				total = total + 1
			}
		}
		return total <= h
	})

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
