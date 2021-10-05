package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 3, 3, 4, 4, 8, 8, 9}))
}

// 剑指OfferII070.排序数组中只出现一次的数字
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
