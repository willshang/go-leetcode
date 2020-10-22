package main

import (
	"fmt"
)

func main() {
	//fmt.Println(increasingTriplet([]int{10, 7, 11, 8, 9, 10}))
	fmt.Println(increasingTriplet([]int{3, 4, 2, 6}))
}

func increasingTriplet(nums []int) bool {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] >= 2 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
