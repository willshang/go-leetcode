package main

import "sort"

func main() {

}

// leetcode2279_装满石头的背包的最大数量
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	res := 0
	for i := 0; i < len(capacity); i++ {
		capacity[i] = capacity[i] - rocks[i]
	}
	sort.Ints(capacity)
	for i := 0; i < len(capacity); i++ {
		if capacity[i] <= additionalRocks {
			res++
			additionalRocks = additionalRocks - capacity[i]
		} else {
			break
		}
	}
	return res
}
