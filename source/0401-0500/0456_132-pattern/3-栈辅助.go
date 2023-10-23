package main

import "math"

func main() {

}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	stack := make([]int, 0)
	maxValue := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		// i < k
		if nums[i] < maxValue {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			last := len(stack) - 1
			maxValue = max(maxValue, stack[last])
			stack = stack[:last]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
