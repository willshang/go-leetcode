package main

import "sort"

func main() {

}

// leetcode2099_找到和最大的长度为K的子序列
func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	arr := make([][2]int, len(nums))
	for i := 0; i < n; i++ {
		arr[i] = [2]int{i, nums[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	sort.Slice(arr[:k], func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, arr[i][1])
	}
	return res
}
