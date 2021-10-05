package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// leetcode40_组合总和II
var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	sort.Ints(candidates)
	dfs(candidates, target, []int{}, 0)
	return res
}

func dfs(candidates []int, target int, arr []int, index int) {
	if target == 0 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		origin := i
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
		arr = append(arr, candidates[i])
		dfs(candidates, target-candidates[i], arr, origin+1)
		arr = arr[:len(arr)-1]
	}
}
