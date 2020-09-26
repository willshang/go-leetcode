package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	// 非负整数数组
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum < int(math.Abs(float64(S))) || (sum+S)%2 == 1 {
		return 0
	}
	// 一个正数x,一个负数背包y => x+y=sum, x-y=S => (sum+S)/2=x
	target := (sum + S) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= len(nums); i++ {
		// 从后往前，避免覆盖
		for j := target; j >= 0; j-- {
			if j >= nums[i-1] {
				// 背包足够大，都选
				dp[j] = dp[j] + dp[j-nums[i-1]]
			} else {
				// 容量不够，不选
				dp[j] = dp[j]
			}
		}
	}
	return dp[target]
}
