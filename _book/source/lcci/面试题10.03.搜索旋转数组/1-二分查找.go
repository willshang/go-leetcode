package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{5, 5, 0, 1, 3, 5}, 5))
}

// 程序员面试金典10.03_搜索旋转数组
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[left] < nums[mid] { // 左边升序的情况
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] { // 右边升序
			if nums[mid] < target && target <= nums[right] && nums[left] > nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else if nums[left] == nums[mid] {
			if nums[left] != target {
				left++
			} else {
				return left
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
