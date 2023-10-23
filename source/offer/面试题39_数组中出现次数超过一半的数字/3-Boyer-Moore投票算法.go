package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

// 剑指offer_面试题39_数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	result, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
		} else if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	return result
}
