package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	count := 0
	if len(nums) > 0 {
		first := getFirstK(nums, target, 0, len(nums)-1)
		last := getLastK(nums, target, 0, len(nums)-1)
		if first > -1 && last > -1 {
			count = last - first + 1
		}
	}
	return count
}

func getFirstK(nums []int, target int, start, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		if (mid > 0 && nums[mid-1] != target) || mid == 0 {
			return mid
		} else {
			end = mid - 1
		}
	} else if nums[mid] > target {
		end = mid - 1
	} else {
		start = mid + 1
	}
	return getFirstK(nums, target, start, end)
}

func getLastK(nums []int, target int, start, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		if (mid < len(nums)-1 && nums[mid+1] != target) || mid == len(nums)-1 {
			return mid
		} else {
			start = mid + 1
		}
	} else if nums[mid] < target {
		start = mid + 1
	} else {
		end = mid - 1
	}
	return getLastK(nums, target, start, end)
}
