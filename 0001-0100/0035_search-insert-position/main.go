package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}

/*
//执行用时：8 ms
func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] < target{
		if nums[i] == target{
			return i
		}
		i++
	}
	return i
}*/

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			return mid
		}
	}
	return low
}
