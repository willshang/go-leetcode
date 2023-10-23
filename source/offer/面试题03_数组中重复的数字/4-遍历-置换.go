package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

// 剑指offer_面试题03_数组中重复的数字
func findRepeatNumber(nums []int) int {
	for key, value := range nums {
		if key == value {
			continue
		}
		if value == nums[value] {
			return nums[value]
		}
		nums[key], nums[value] = nums[value], nums[key]
	}
	return -1
}
