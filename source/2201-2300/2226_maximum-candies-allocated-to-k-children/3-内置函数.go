package main

import "sort"

func main() {

}

func maximumCandies(candies []int, k int64) int {
	sum := int64(0)
	maxValue := 0
	for i := 0; i < len(candies); i++ {
		sum = sum + int64(candies[i])
		maxValue = max(maxValue, candies[i])
	}
	if sum < k {
		return 0
	}
	return sort.Search(maxValue+1, func(target int) bool {
		if target == 0 {
			return false
		}
		res := int64(0)
		for i := 0; i < len(candies); i++ {
			res = res + int64(candies[i]/target)
		}
		return res < k
	}) - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
