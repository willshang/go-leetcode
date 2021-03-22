package main

import "sort"

func main() {

}

// leetcode1798_你能构造出连续值的最大数目
func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	res := 0
	for i := 0; i < len(coins); i++ {
		if res >= coins[i]-1 {
			res = res + coins[i]
		} else {
			break
		}
	}
	return res + 1
}
