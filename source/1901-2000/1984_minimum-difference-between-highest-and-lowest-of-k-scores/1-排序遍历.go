package main

import (
	"math"
	"sort"
)

func main() {

}

// leetcode1984_学生分数的最小差值
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	res := math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		left := nums[i]
		right := nums[i+k-1]
		res = min(res, right-left)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
