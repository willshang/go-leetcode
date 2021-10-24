package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 9, 7}, 4))
}

// leetcode1438_绝对差不超过限制的最长连续子数组
func longestSubarray(nums []int, limit int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		maxValue, minValue := nums[i], nums[i]
		for j := i; j < len(nums); j++ {
			maxValue = max(maxValue, nums[j])
			minValue = min(minValue, nums[j])
			if maxValue-minValue <= limit {
				res = max(res, j-i+1)
			} else {
				break
			}
		}
		if res >= len(nums)-i {
			break
		}
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
