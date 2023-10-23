package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

// 剑指OfferII101.分割等和子集
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	// 题目转换为0-1背包问题，容量为sum/2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= 0; j-- {
			if j-nums[i] >= 0 && dp[j-nums[i]] == true {
				dp[j] = true
			}
		}
	}
	return dp[target]
}
