package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	sort.Ints(candidates)
	dfs(candidates, target, []int{}, 0)
	return res
}
