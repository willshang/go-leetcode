package main

import "fmt"

func main() {
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{1, 3, 5}, 1))
}

// leetcode81_搜索旋转排序数组II
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
