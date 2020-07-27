package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{4, 7, 9, 0, 3, 4, 2, 1}
	arr := []int{1, 3, 2}
	// [4 7 9 0 3 4 2 1]
	// [4 7 9 0 4 3 2 1]
	// [4 7 9 0 4 1 2 3]
	nextPermutation(arr)
	fmt.Println(arr)
}

// leetcode31_下一个排列
func nextPermutation(nums []int) {
	left := len(nums) - 2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left == -1 {
		sort.Ints(nums)
		return
	}
	right := len(nums) - 1
	for right >= 0 && nums[right] <= nums[left] {
		right--
	}
	nums[left], nums[right] = nums[right], nums[left]
	// sort.Ints(nums[left+1:])
	count := 0
	for i := left + 1; i <= (left+1+len(nums)-1)/2; i++ {
		nums[i], nums[len(nums)-1-count] = nums[len(nums)-1-count], nums[i]
		count++
	}
}
