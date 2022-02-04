package main

import "sort"

func main() {

}

func minOperations(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	// 去重
	index := 0
	for i := 1; i < n; i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	nums = nums[:index+1]
	// 二分查找
	res := 0
	for right := 0; right < len(nums); right++ {
		// 计算区间[nums[i]-n+1, nums[i]]的长度
		left := sort.SearchInts(nums[:right], nums[right]-n+1)
		res = max(res, right-left+1) // 取最长的长度
	}
	return n - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
