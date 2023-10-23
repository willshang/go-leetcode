package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := left
	for nums[left] >= nums[right] {
		if right-left == 1 {
			mid = right
			break
		}
		mid = (left + right) / 2
		if nums[left] == nums[right] && nums[mid] == nums[left] {
			return minInorder(nums, left, right)
		}
		if nums[mid] >= nums[left] {
			left = mid
		} else if nums[mid] <= nums[right] {
			right = mid
		}
	}
	return nums[mid]
}

func minInorder(numbers []int, left, right int) int {
	result := numbers[left]
	for i := left + 1; i <= right; i++ {
		if result > numbers[i] {
			result = numbers[i]
		}
	}
	return result
}
