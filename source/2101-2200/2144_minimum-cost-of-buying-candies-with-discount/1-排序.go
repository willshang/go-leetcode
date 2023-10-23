package main

import "sort"

func main() {

}

// leetcode2144_打折购买糖果的最小开销
func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})
	res := 0
	for i := 0; i < len(cost); i++ {
		if i%3 != 2 {
			res = res + cost[i]
		}
	}
	return res
}
