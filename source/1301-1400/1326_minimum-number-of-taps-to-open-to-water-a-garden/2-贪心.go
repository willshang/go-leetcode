package main

import "fmt"

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
}

func minTaps(n int, ranges []int) int {
	// 题目变成leetcode45.跳跃游戏II
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		l := max(0, i-ranges[i])
		r := min(n, i+ranges[i])
		arr[l] = max(arr[l], r) // 更新当前位置能到达最远的位置
	}
	last := 0
	prev := 0
	res := 0
	// 变成leetcode45.跳跃游戏II的变形
	for i := 0; i < len(arr); i++ {
		if arr[i] > last {
			last = arr[i]
		}
		if i == last && last < n { // 无法达到目标
			return -1
		}
		if i == prev && i < n {
			res++
			prev = last
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
