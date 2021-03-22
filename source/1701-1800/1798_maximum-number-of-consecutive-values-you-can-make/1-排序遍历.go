package main

import "sort"

func main() {

}

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	res := 1
	target := 1
	sum := 0
	for i := 0; i < len(coins); i++ {
		sum = sum + coins[i]
		if coins[i] <= target {
			target = sum + 1
			res = target
		} else {
			break
		}
	}
	return res
}
