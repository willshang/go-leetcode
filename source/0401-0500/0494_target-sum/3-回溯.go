package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

var res int

func findTargetSumWays(nums []int, S int) int {
	res = 0
	dfs(nums, 0, S)
	return res
}

func dfs(nums []int, index int, target int) {
	if index == len(nums) {
		if target == 0 {
			res++
		}
		return
	}
	target = target + nums[index]
	dfs(nums, index+1, target)
	target = target - nums[index]
	//
	target = target - nums[index]
	dfs(nums, index+1, target)
	target = target + nums[index]
}
