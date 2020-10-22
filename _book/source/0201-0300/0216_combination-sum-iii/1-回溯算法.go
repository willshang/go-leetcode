package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

// leetcode216_组合总和III
var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	arr := make([]int, 0)
	dfs(k, n, 1, arr)
	return res
}

func dfs(k, n int, level int, arr []int) {
	if k == 0 || n < 0 {
		if n == 0 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			res = append(res, temp)
		}
		return
	}
	for i := level; i <= 9; i++ {
		dfs(k-1, n-i, i+1, append(arr, i))
	}
}
