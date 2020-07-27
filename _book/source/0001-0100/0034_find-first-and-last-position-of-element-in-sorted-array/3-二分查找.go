package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	//fmt.Println(searchRange([]int{1}, 0))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	left := leftSearch(nums, target)
	right := rightSearch(nums, target)
	return []int{left, right}
}

func leftSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func rightSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1
}
