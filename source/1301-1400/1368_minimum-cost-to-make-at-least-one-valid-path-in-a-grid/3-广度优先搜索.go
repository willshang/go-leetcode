package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCost([][]int{
		{1, 1, 3},
		{3, 2, 2},
		{1, 1, 4},
	}))
}

// leetcode1368_使网格图至少有一条有效路径的最小代价
// 右左下上
var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func minCost(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	total := n * m
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arr[i][j] = total
		}
	}
	arr[0][0] = 0
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		a, b := node[0], node[1]
		for i := 0; i < 4; i++ {
			x, y := a+dx[i], b+dy[i]
			if 0 <= x && x < n && 0 <= y && y < m {
				dis := arr[a][b]
				if i+1 != grid[a][b] {
					dis = dis + 1
				}
				if dis < arr[x][y] {
					arr[x][y] = dis
					queue = append(queue, []int{x, y})
				}
			}
		}
	}
	return arr[n-1][m-1]
}
