package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	dfs(n, k, 1, make([]int, 0))
	return res
}

func dfs(n, k, index int, arr []int) {
	if len(arr) == k {
		temp := make([]int, k)
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := index; i <= n; i++ {
		arr = append(arr, i)
		dfs(n, k, i+1, arr)
		arr = arr[:len(arr)-1]
	}
}
