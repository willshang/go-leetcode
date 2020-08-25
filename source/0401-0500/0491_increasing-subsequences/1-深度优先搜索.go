package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

// leetcode491_递增子序列
var res [][]int

func findSubsequences(nums []int) [][]int {
	res = make([][]int, 0)
	dfs(nums, 0, math.MinInt32, make([]int, 0))
	return res
}

func dfs(nums []int, index int, prev int, arr []int) {
	if index == len(nums) {
		if len(arr) >= 2 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			res = append(res, temp)
		}
		return
	}
	// 选nums[index]
	if prev <= nums[index] {
		arr = append(arr, nums[index])
		dfs(nums, index+1, nums[index], arr)
		arr = arr[:len(arr)-1]
	}
	// 不选nums[index]
	if prev != nums[index] {
		dfs(nums, index+1, prev, arr)
	}
}
