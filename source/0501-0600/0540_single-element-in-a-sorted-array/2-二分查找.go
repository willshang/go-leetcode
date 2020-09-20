package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 3, 3, 4, 4, 8, 8, 9}))
}

// leetcode540_有序数组中的单一元素
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			right = mid
		}
	}
	return nums[left]
}
