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
	total := 1 << n
	arr := make([]int, total)
	dp := make([]bool, total)
	dp[0] = true
	for i := 0; i < total; i++ { // 枚举状态
		if dp[i] == false {
			continue
		}
		for j := 0; j < n; j++ { // 基于当前状态，添加1个数
			if i&(1<<j) > 0 { // 第j位为1，跳过
				continue
			}
			next := i | (1 << j) // 添加完后的值
			if dp[next] == true {
				continue
			}
			if arr[i]+nums[j] <= target {
				arr[next] = (arr[i] + nums[j]) % target
				dp[next] = true
			} else {
				break // 已经排好序，后面会更大
			}
		}
	}
	return dp[total-1]
}
