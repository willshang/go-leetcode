package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4}))
}

func maxAbsoluteSum(nums []int) int {
	arr := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	maxValue, minValue := 0, 0
	res := 0
	for i := 1; i < len(arr); i++ {
		res = max(res, abs(arr[i]-maxValue))
		res = max(res, abs(arr[i]-minValue))
		maxValue = max(maxValue, arr[i])
		minValue = min(minValue, arr[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
