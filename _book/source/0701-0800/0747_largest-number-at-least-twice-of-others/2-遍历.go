package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
	fmt.Println(dominantIndex([]int{1, 0}))
}

// leetcode747_至少是其他数字两倍的最大数
func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	maxValue := nums[0]
	index := 0

	for i := 1; i < n; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			index = i
		}
	}
	for i := 0; i < n; i++ {
		if i == index {
			continue
		}
		if nums[i]*2 > maxValue {
			return -1
		}
	}
	return index
}
