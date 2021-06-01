package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{4, 7, 9, 0, 3, 4, 2, 1}
	//arr := []int{1, 3, 2}
	arr := []int{6, 9, 7, 5}
	// [4 7 9 0 3 4 2 1]
	// [4 7 9 0 4 3 2 1]
	// [4 7 9 0 4 1 2 3]
	for i := 0; i < 8; i++ {
		nextPermutation(arr)
		fmt.Println(arr)
	}
}

// leetcode31_下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	left := n - 2
	// 以12385764为例，从后往前找到5<7 的升序情况，目标值为左边的数5
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left == -1 {
		sort.Ints(nums)
		return
	}
	right := n - 1
	// 从后往前，找到第一个大于目标值的数，如6>5，然后交换
	for right >= 0 && nums[right] <= nums[left] {
		right--
	}
	nums[left], nums[right] = nums[right], nums[left]
	count := 0
	// 后面是降序状态，让它变为升序
	for i := left + 1; i <= (left+1+n-1)/2; i++ {
		nums[i], nums[n-1-count] = nums[n-1-count], nums[i]
		count++
	}
}
