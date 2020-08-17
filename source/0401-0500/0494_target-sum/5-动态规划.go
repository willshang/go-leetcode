package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

// leetcode494_目标和
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
	// 在前i个物品中选择，若当前背包的容量为j，则最多有x种方法可以恰好装满背包。
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1 // 容量为0， 只有都不选
	}
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i-1] {
				// 背包足够大，都选
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				// 容量不够，不选
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][target]
}
