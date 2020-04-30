package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(arr))
}

func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	left, right := 0, -1
	min, max := nums[length-1], nums[0]
	for i := 1; i < length; i++ {
		fmt.Println(i, max, nums[i], right)
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}

		j := length - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			left = j
		}
	}
	return right - left + 1
}
