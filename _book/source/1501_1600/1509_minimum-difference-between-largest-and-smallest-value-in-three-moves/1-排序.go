package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDifference([]int{5, 3, 2, 4}))
}

// leetcode1509_三次操作后最大值与最小值的最小差
func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	sort.Ints(nums)
	res := math.MaxInt32
	for i := 0; i <= 3; i++ {
		res = min(res, nums[len(nums)-1-(3-i)]-nums[i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
