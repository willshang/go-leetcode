package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
