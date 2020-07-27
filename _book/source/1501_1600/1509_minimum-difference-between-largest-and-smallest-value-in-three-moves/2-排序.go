package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDifference([]int{5, 3, 2, 4}))
}

func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	sort.Ints(nums)
	res := math.MaxInt32
	res = min(res, nums[len(nums)-1]-nums[3])
	res = min(res, nums[len(nums)-2]-nums[2])
	res = min(res, nums[len(nums)-3]-nums[1])
	res = min(res, nums[len(nums)-4]-nums[0])
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
