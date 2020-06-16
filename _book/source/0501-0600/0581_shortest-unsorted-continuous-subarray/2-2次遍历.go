package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	//arr := []int{2, 3, 4, 5, 6, 9, 15}
	fmt.Println(findUnsortedSubarray(arr))
}

func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	right := -1
	max := nums[0]
	for i := 1; i < length; i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}
	}
	if right == 0 {
		// 针对升序，特殊处理
		// 如去掉判断
		// 需要保证left,right初始值满足right-left+1=0
		return 0
	}
	left := 0
	min := nums[length-1]
	for i := length - 2; i >= 0; i-- {
		if min >= nums[i] {
			min = nums[i]
		} else {
			left = i
		}
	}
	return right - left + 1
}
