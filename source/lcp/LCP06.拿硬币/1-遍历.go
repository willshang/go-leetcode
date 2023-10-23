package main

import "fmt"

func main() {
	fmt.Println(minCount([]int{4, 2, 1}))
}

// leetcode_lcp06_拿硬币
func minCount(coins []int) int {
	res := 0
	for i := 0; i < len(coins); i++ {
		res = res + coins[i]/2
		if coins[i]%2 == 1 {
			res = res + 1
		}
	}
	return res
}
