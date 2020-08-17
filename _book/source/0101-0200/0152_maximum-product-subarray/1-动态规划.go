package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	minValue, maxValue, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minV, maxV := minValue, maxValue
		minValue = min(minV*nums[i], min(nums[i], maxV*nums[i]))
		maxValue = max(maxV*nums[i], max(nums[i], minV*nums[i]))
		res = max(res, maxValue)
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
