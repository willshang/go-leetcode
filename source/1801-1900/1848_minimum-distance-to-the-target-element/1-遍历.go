package main

import "math"

func main() {

}

// leetcode1848_到目标元素的最小距离
func getMinDistance(nums []int, target int, start int) int {
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = min(res, abs(i-start))
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
