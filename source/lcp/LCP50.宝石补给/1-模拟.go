package main

import "math"

func main() {

}

// leetcode_lcp50_宝石补给
func giveGem(gem []int, operations [][]int) int {
	for i := 0; i < len(operations); i++ {
		a, b := operations[i][0], operations[i][1]
		v := gem[a] / 2
		gem[a] = gem[a] - v
		gem[b] = gem[b] + v
	}
	maxValue, minValue := math.MinInt32, math.MaxInt32
	for i := 0; i < len(gem); i++ {
		maxValue = max(maxValue, gem[i])
		minValue = min(minValue, gem[i])
	}
	return maxValue - minValue
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
