package main

import "fmt"

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
}

// leetcode1262_可被三整除的最大和
func maxSumDivThree(nums []int) int {
	dp := [3]int{}
	for i := 0; i < len(nums); i++ {
		for _, v := range dp {
			value := v + nums[i]
			dp[value%3] = max(dp[value%3], value)
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
