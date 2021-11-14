package main

import "fmt"

func main() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
}

var n int
var maxValue int
var res int
var left, right []int

func countHighestScoreNodes(parents []int) int {
	res = 0
	maxValue = 1
	n = len(parents)
	left, right = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		left[i], right[i] = -1, -1
	}
	for i := 1; i < n; i++ { // 建立左右子树关系
		a, b := i, parents[i] // b=>a
		if left[b] == -1 {    // 优先放在左子树
			left[b] = a
		} else {
			right[b] = a
		}
	}
	dfs(0)
	return res
}

func dfs(root int) int {
	if root == -1 {
		return 0
	}
	a, b := dfs(left[root]), dfs(right[root])
	value := max(a, 1) * max(b, 1) * max(n-a-b-1, 1)
	if value > maxValue {
		maxValue = value
		res = 1
	} else if value == maxValue {
		res++
	}
	return a + b + 1 // 左子树+右子树+1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
