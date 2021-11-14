package main

import "fmt"

func main() {
	fmt.Println(maximalPathQuality([]int{5, 10, 15, 20}, [][]int{
		{0, 1, 10},
		{1, 2, 10},
		{0, 3, 10},
	}, 30))
}

// leetcode2065_最大化一张图中的路径价值
var res int
var arr [][][2]int
var visited []int

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	arr = make([][][2]int, n) // 邻接表
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1], edges[i][2]
		arr[a] = append(arr[a], [2]int{b, c})
		arr[b] = append(arr[b], [2]int{a, c})
	}
	res = 0
	visited = make([]int, n)
	dfs(values, maxTime, 0, 0, 0)
	return res
}

func dfs(values []int, maxTime int, start int, t int, sum int) {
	if t > maxTime {
		return
	}

	if visited[start] == 0 {
		sum = sum + values[start]
	}
	visited[start]++
	if start == 0 {
		res = max(res, sum)
	}
	for i := 0; i < len(arr[start]); i++ {
		next, c := arr[start][i][0], arr[start][i][1]
		dfs(values, maxTime, next, t+c, sum)
	}
	visited[start]--
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
