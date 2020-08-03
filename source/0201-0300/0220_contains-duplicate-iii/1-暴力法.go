package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}

// leetcode220_存在重复元素III
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if abs(nums[i], nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
