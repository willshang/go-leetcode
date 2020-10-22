package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

func canPartition(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	return dfs(nums, target, 0)
}

func dfs(nums []int, target int, index int) bool {
	if target == 0 {
		return true
	}
	for i := index; i < len(nums); i++ {
		if index < i && nums[i] == nums[i-1] {
			continue
		}
		if target-nums[i] < 0 {
			return false
		}
		if dfs(nums, target-nums[i], i+1) == true {
			return true
		}
	}
	return false
}
