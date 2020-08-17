package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{1}, 0))
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] != target {
			if target > nums[left] {
				left++
			}
		}
		if nums[right] != target {
			if target < nums[right] {
				right--
			}
		}
		if left < len(nums) && right >= 0 &&
			nums[left] == nums[right] && nums[left] == target {
			break
		}
	}
	if right < left {
		return []int{-1, -1}
	}
	return []int{left, right}
}
