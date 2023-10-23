package main

import (
	"math"
	"sort"
)

func main() {

}

func findLonely(nums []int) []int {
	res := make([]int, 0)
	nums = append(nums, math.MinInt32)
	nums = append(nums, math.MaxInt32)
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i++ {
		if nums[i]-nums[i-1] > 1 && nums[i+1]-nums[i] > 1 {
			res = append(res, nums[i])
		}
	}
	return res
}
