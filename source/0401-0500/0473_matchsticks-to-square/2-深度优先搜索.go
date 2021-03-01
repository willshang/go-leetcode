package main

import "sort"

func main() {

}

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
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	visited := make([]bool, len(nums))
	for i := 0; i < 4; i++ {
		if dfs(nums, sum/4, 0, 0, visited) == false {
			return false
		}
	}
	return true
}

func dfs(nums []int, target int, sum int, level int, visited []bool) bool {
	if sum == target {
		return true
	}
	if len(nums) == level {

	}
	for i := level; i < len(nums); i++ {
		if visited[i] == false && sum <= target {
			visited[i] = true
			if dfs(nums, target, sum+nums[i], level+1, visited) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}
