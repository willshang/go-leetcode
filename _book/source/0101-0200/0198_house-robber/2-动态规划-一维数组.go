package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

/*
State:      dp[i]，表示到第i个房子时能够抢到的最大金额。
Function:   dp[i] = max(num[i] + dp[i - 2], dp[i - 1])
Initialize: dp[0] = num[0], dp[1] = max(num[0], num[1])
Return:     dp[n]
*/
// 保存一维dp数组
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
