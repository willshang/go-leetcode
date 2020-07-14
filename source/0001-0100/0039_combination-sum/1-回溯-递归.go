package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
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
		if target < candidates[i] {
			return
		}
		// dfs(candidates, target-candidates[i], append(arr, candidates[i]), i)
		arr = append(arr, candidates[i])
		dfs(candidates, target-candidates[i], arr, i)
		arr = arr[:len(arr)-1]
	}
}
