package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
	fmt.Println(dominantIndex([]int{1, 0}))
}

func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	maxIndex, secondMaxIndex := 0, 1
	if nums[maxIndex] < nums[secondMaxIndex] {
		maxIndex, secondMaxIndex = secondMaxIndex, maxIndex
	}

	for i := 2; i < n; i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex, secondMaxIndex = i, maxIndex
		} else if nums[secondMaxIndex] < nums[i] {
			secondMaxIndex = i
		}
	}
	if nums[maxIndex] >= 2*nums[secondMaxIndex] {
		return maxIndex
	}
	return -1
}
