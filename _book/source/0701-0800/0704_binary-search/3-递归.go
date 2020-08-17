package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return search(nums[:mid], target)
	} else {
		result := search(nums[mid+1:], target)
		if result == -1 {
			return result
		}
		return mid + 1 + result
	}
}
