package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(arr))
}

// leetcode217_存在重复元素
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
