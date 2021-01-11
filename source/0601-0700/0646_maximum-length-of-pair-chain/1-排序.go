package main

import (
	"math"
	"sort"
)

func main() {

}

// leetcode646_最长数对链
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	res := 0
	cur := math.MinInt32
	for i := 0; i < len(pairs); i++ {
		if cur < pairs[i][0] {
			cur = pairs[i][1]
			res++
		}
	}
	return res
}
