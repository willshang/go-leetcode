package main

import (
	"sort"
)

func main() {

}

func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	n := len(nums)
	sort.Ints(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	if sum%k != 0 { // 分不开：false
		return false
	}
	target := sum / k
	if nums[n-1] > target { // 有1个大于平均值：false
		return false
	}
	return dfs(nums, k, target, 0, 0, make([]bool, n))
}

func dfs(nums []int, k int, target int, index int, count int, visited []bool) bool {
	if k == 0 {
		return true
	}
	if count == target {
		return dfs(nums, k-1, target, 0, 0, visited) // k减少一个，从头开始
	}
	for i := index; i < len(nums); i++ {
		if visited[i] == true { // nums[i]使用过
			continue
		}
		if count+nums[i] > target { // 大于目标值
			continue
		}
		visited[i] = true
		count = count + nums[i]
		if dfs(nums, k, target, i+1, count, visited) == true {
			return true
		}
		count = count - nums[i]
		visited[i] = false
	}
	return false
}
