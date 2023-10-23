package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

// 剑指OfferII080.含有k个元素的组合
var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	dfs(n, k, 1, make([]int, 0))
	return res
}

func dfs(n, k, index int, arr []int) {
	if index > n+1 {
		return
	}
	if len(arr) == k {
		temp := make([]int, k)
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	dfs(n, k, index+1, arr)
	arr = append(arr, index)
	dfs(n, k, index+1, arr)
}
