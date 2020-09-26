package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}

// leetcode162_寻找峰值
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	left := 0
	right := n - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
