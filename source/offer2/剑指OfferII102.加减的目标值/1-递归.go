package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == 0 && S == 0 {
			return 2
		}
		if nums[0] == S || nums[0] == -S {
			return 1
		}
	}
	value := nums[0]
	nums = nums[1:]
	return findTargetSumWays(nums, S-value) + findTargetSumWays(nums, S+value)
}
