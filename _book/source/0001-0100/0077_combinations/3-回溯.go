package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

// leetcode77_组合
var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	dfs(nums, 0, k, make([]int, 0))
	return res
}

func dfs(nums []int, index, k int, arr []int) {
	if len(arr) == k {
		temp := make([]int, k)
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		arr = append(arr, nums[i])
		dfs(nums, i+1, k, arr)
		arr = arr[:len(arr)-1]
	}
}
