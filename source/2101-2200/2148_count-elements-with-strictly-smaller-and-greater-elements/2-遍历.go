package main

import "fmt"

func main() {
	fmt.Println(countElements([]int{-3, 3, 3, 90}))
}

// leetcode2148_元素计数
func countElements(nums []int) int {
	minValue, maxValue := nums[0], nums[0]
	minCount, maxCount := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxCount = 1
		} else if nums[i] == maxValue {
			maxCount++
		}
		if nums[i] < minValue {
			minValue = nums[i]
			minCount = 1
		} else if nums[i] == minValue {
			minCount++
		}
	}
	if maxValue == minValue {
		return 0
	}
	return len(nums) - maxCount - minCount
}
