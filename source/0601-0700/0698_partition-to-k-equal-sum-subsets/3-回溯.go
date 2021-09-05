package main

import (
	"sort"
)

func main() {

}

// leetcode698_划分为k个相等的子集
func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	if sum%k != 0 { // 分不开：false
		return false
	}
	target := sum / k
	if nums[0] > target { // 有1个大于平均值：false
		return false
	}
	return dfs(nums, k, target, 0, make([]int, k))
}

// 不做剪枝，需要排序从大到小
func dfs(nums []int, k int, target int, index int, sum []int) bool {
	if index == len(nums) {
		return true
	}
	for i := 0; i < k; i++ {
		if sum[i] < target && sum[i]+nums[index] <= target {
			sum[i] = sum[i] + nums[index]
			if dfs(nums, k, target, index+1, sum) == true {
				return true
			}
			sum[i] = sum[i] - nums[index]
		}
	}
	return false
}
