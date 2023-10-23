package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}

// leetcode135_分发糖果
func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	right := make([]int, n)
	// 规则1：每个孩子至少分配到 1 个糖果。
	for i := 0; i < n; i++ {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = res + max(left[i], right[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
