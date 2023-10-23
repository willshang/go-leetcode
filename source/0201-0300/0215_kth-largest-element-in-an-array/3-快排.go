package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

// leetcode215_数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	return findK(nums, 0, len(nums)-1, k)
}

func findK(nums []int, start, end int, k int) int {
	if start >= end {
		return nums[end]
	}
	index := partition(nums, start, end)
	if index+1 == k {
		return nums[index]
	} else if index+1 < k {
		return findK(nums, index+1, end, k)
	}
	return findK(nums, start, index-1, k)
}

func partition(nums []int, start, end int) int {
	temp := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] > temp {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
