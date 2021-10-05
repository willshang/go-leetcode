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
// 不保存dp数组
// 剑指OfferII089.房屋偷盗
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := nums[0]
	b := max(a, nums[1])

	for i := 2; i < len(nums); i++ {
		a, b = b, max(a+nums[i], b)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
