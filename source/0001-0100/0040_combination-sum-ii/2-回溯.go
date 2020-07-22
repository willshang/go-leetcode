package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

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
	for i := index; i < len(candidates); i++ {
		if i != index && candidates[i] == candidates[i-1] {
			continue
		}
		if target < 0 {
			return
		}
		arr = append(arr, candidates[i])
		dfs(candidates, target-candidates[i], arr, i+1)
		arr = arr[:len(arr)-1]
	}
}
