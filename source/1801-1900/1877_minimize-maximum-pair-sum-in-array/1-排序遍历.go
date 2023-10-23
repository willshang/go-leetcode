package main

import "sort"

func main() {

}

// leetcode1877_数组中最大数对和的最小值
func minPairSum(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	res := 0
	for i := 0; i < n/2; i++ {
		res = max(res, nums[i]+nums[n-1-i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
