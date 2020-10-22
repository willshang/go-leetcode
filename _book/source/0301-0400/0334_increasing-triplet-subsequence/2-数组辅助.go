package main

import (
	"fmt"
)

func main() {
	//fmt.Println(increasingTriplet([]int{10, 7, 11, 8, 9, 10}))
	fmt.Println(increasingTriplet([]int{3, 4, 2, 6}))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	a := make([]int, len(nums))
	b := make([]int, len(nums))
	a[0] = nums[0]
	b[len(b)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		a[i] = min(a[i-1], nums[i])
	}
	for i := len(nums) - 2; i >= 0; i-- {
		b[i] = max(b[i+1], nums[i])
	}
	for i := 0; i < len(nums); i++ {
		if a[i] < nums[i] && nums[i] < b[i] {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
