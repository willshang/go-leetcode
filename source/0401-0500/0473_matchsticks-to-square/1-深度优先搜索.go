package main

import "sort"

func main() {

}

// leetcode473_火柴拼正方形
func makesquare(nums []int) bool {
	n := len(nums)
	if nums == nil || n < 4 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum%4 != 0 {
		return false
	}
	// 从大到小排序；从小到大可能会超时
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	res := make([]int, 4)
	return dfs(nums, res, sum/4, 0)
}

func dfs(nums []int, res []int, target int, level int) bool {
	if len(nums) == level {
		for i := 0; i < len(res); i++ {
			if res[i] != target {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(res); i++ {
		if nums[level]+res[i] > target {
			continue
		}
		res[i] = res[i] + nums[level]
		if dfs(nums, res, target, level+1) == true {
			return true
		}
		res[i] = res[i] - nums[level]
	}
	return false
}
