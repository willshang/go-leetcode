package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4}))
}

func maxAbsoluteSum(nums []int) int {
	maxValue, minValue := 0, 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if maxValue >= 0 {
			maxValue = maxValue + nums[i]
		} else {
			maxValue = nums[i]
		}
		if minValue <= 0 {
			minValue = minValue + nums[i]
		} else {
			minValue = nums[i]
		}
		res = max(res, abs(maxValue))
		res = max(res, abs(minValue))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
