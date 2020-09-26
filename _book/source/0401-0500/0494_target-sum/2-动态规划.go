package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {
	dp := make(map[int]int)
	dp[nums[0]]++
	dp[-nums[0]]++
	for i := 1; i < len(nums); i++ {
		temp := make(map[int]int)
		for k, v := range dp {
			temp[k-nums[i]] = temp[k-nums[i]] + v
			temp[k+nums[i]] = temp[k+nums[i]] + v
		}
		dp = temp
	}
	return dp[S]
}
