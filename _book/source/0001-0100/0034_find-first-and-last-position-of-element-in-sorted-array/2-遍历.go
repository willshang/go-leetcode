package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{1}, 0))
}

func searchRange(nums []int, target int) []int {
	left := -1
	right := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			right = i
		} else if nums[i] > target {
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			left = i
		} else if nums[i] < target {
			break
		}
	}
	return []int{left, right}
}
