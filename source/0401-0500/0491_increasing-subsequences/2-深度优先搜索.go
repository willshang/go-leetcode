package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

var res [][]int

func findSubsequences(nums []int) [][]int {
	res = make([][]int, 0)
	dfs(nums, 0, make([]int, 0))
	return res
}

func dfs(nums []int, index int, arr []int) {
	if len(arr) >= 2 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
	}
	m := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if m[nums[i]] == true || (len(arr) > 0 && nums[i] < arr[len(arr)-1]) {
			continue
		}
		m[nums[i]] = true
		dfs(nums, i+1, append(arr, nums[i]))
	}
}
