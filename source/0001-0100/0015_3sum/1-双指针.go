package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// leetcode15_三数之和
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		target := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		if nums[i] > 0 || nums[i]+nums[left] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-2 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right] > target {
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return res
}
