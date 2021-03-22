package main

import (
	"fmt"
)

func main() {
	fmt.Println(check([]int{2, 1, 3, 4}))
}

// leetcode1752_检查数组是否经排序和轮转得到
func check(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if nums[0] < nums[len(nums)-1] {
				return false
			}
		}
	}
	return true
}
