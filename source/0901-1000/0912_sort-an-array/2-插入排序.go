package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, 1}
	sortArray(arr)
	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		pos := i - 1 // 已经有序的最后一个下标
		cur := nums[i]
		for pos >= 0 && nums[pos] > cur {
			nums[pos+1] = nums[pos] // 后移
			pos--
		}
		nums[pos+1] = cur
	}
	return nums
}
