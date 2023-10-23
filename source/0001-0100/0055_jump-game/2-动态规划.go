package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		flag := false
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				flag = true
				break
			}
		}
		dp[i] = flag
	}
	return dp[len(nums)-1]
}
