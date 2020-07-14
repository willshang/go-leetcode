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
	dfs(nums, 0, make([]int, len(nums)), make([]int, 0))
	return res
}

func dfs(nums []int, index int, visited []int, arr []int) {
	if len(nums) == index {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 1 {
			continue
		}
		arr = append(arr, nums[i])
		visited[i] = 1
		dfs(nums, index+1, visited, arr)
		visited[i] = 0
		arr = arr[:len(arr)-1]
	}
}
