package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

// leetcode78_子集
var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	dfs(nums, make([]int, 0), 0)
	return res
}

func dfs(nums []int, arr []int, level int) {
	if level >= len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	dfs(nums, arr, level+1)
	dfs(nums, append(arr, nums[level]), level+1)
}
