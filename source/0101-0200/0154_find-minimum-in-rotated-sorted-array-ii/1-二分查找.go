package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

// leetcode154_寻找旋转排序数组中的最小值II
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
