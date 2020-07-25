package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1, 4, 4, 4, 4}))
}

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res = make([][]int, 0)
	dfs(nums, make([]int, 0), 0)
	return res
}

func dfs(nums []int, arr []int, level int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	res = append(res, temp)
	for i := level; i < len(nums); i++ {
		if i > level && nums[i] == nums[i-1] {
			continue
		}
		arr = append(arr, nums[i])
		dfs(nums, arr, i+1)
		arr = arr[:len(arr)-1]
	}
}
