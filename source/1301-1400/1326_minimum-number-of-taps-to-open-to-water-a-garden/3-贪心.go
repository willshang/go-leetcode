package main

import "fmt"

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
}

func minTaps(n int, ranges []int) int {
	// 处理成一组start-end的数组
	// 题目变成leetcode45.跳跃游戏II
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		l := max(0, i-ranges[i])
		r := min(n, i+ranges[i])
		for j := l; j < r; j++ {
			arr[j] = max(arr[j], r) // 更新当前位置能到达最远的位置
		}
	}
	res := 0
	cur := 0
	for cur < n {
		if arr[cur] == 0 {
			return -1
		}
		cur = arr[cur]
		res++
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
