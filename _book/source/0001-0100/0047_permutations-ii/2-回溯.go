package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, index int) {
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		res = append(res, temp)
		return
	}
	m := make(map[int]int)
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = 1
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
