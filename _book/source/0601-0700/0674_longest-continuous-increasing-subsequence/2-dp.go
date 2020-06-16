package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(arr))
}

// 状态转移方程
// 若nums[i-1]<nums[i]，则dp[i]=dp[i-1]+1；否则dp[i]=1
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
