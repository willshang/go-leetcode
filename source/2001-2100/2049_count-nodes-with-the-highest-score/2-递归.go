package main

import "fmt"

func main() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
}

// leetcode2049_统计最高分的节点数目
var arr [][]int
var n int
var maxValue int
var res int

func countHighestScoreNodes(parents []int) int {
	res = 0
	maxValue = 1
	n = len(parents)
	arr = make([][]int, n)
	for i := 1; i < n; i++ { // 邻接表
		a, b := i, parents[i] // b=>a
		arr[b] = append(arr[b], a)
	}
	dfs(0)
	return res
}

func dfs(root int) int {
	count := 1
	value := 1
	for i := 0; i < len(arr[root]); i++ { // 计算左右子树乘积
		size := dfs(arr[root][i])
		count = count + size
		value = value * size
	}
	if root > 0 { // 计算去掉以根节点位i的剩余部分
		value = value * (n - count)
	}
	if value > maxValue {
		maxValue = value
		res = 1
	} else if value == maxValue {
		res++
	}
	return count
}
