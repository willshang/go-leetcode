package main

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
}

func PredictTheWinner(nums []int) bool {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[len(nums)-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
