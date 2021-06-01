package main

import "fmt"

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// L~R范围的组合数=0~R范围的组合数- 0~L-1范围的组合数
	return count(nums, right) - count(nums, left-1)
}

func count(nums []int, target int) int {
	res := 0
	n := len(nums)
	dp := make([]int, n)
	if nums[0] <= target {
		dp[0] = 1
	}
	res = res + dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= target {
			dp[i] = dp[i-1] + 1
		}
		res = res + dp[i]
	}
	return res
}
