package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}

// 剑指OfferII057.值和下标之差都在给定的范围内
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
