package main

import "sort"

func main() {

}

// leetcode368_最大整除子集
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	sort.Ints(nums)
	dp := make([][]int, n)
	dp[0] = append(dp[0], nums[0])
	resLength := 0
	index := 0
	for i := 1; i < n; i++ {
		maxLength := 0
		arr := make([]int, 0)
		for j := 0; j < i; j++ {
			// 可除以整除集合中的最大值=>属于该集合
			if nums[i]%nums[j] == 0 && len(dp[j]) >= maxLength {
				maxLength = len(dp[j])
				arr = dp[j]
			}
		}
		dp[i] = append(dp[i], append(arr, nums[i])...)
		if len(dp[i]) > resLength {
			resLength = len(dp[i])
			index = i
		}
	}
	return dp[index]
}
