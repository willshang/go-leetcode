package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

// leetcode55_跳跃游戏
func canJump(nums []int) bool {
	j := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= j {
			j = i
		}
	}
	return j <= 0
}
