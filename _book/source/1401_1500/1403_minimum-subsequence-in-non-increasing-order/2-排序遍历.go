package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))
}

// leetcode1403_非递增顺序的最小子序列
func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	target := sum / 2
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum > target {
			return nums[:i+1]
		}
	}
	return nil
}
