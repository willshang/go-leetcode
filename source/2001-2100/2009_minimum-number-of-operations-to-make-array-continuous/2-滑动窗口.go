package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minOperations([]int{4, 2, 5, 3}))
}

// leetcode2009_使数组连续的最少操作数
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
	// 滑动窗口
	res := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > n-1 {
			left++
		}
		res = max(res, right-left+1)
	}
	return n - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
