package main

import "fmt"

func main() {
	fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))
}

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1) // dp[i]表示长度为i的数组nums[:i]能否划分成功
	dp[0] = true
	for i := 0; i < n; i++ {
		// 3种有效划分方式
		if i > 0 && dp[i-1] == true && nums[i] == nums[i-1] {
			dp[i+1] = true
		}
		if i > 1 && dp[i-2] == true && nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			dp[i+1] = true
		}
		if i > 1 && dp[i-2] == true && nums[i] == nums[i-1]+1 && nums[i] == nums[i-2]+2 {
			dp[i+1] = true
		}
	}
	return dp[n]
}
