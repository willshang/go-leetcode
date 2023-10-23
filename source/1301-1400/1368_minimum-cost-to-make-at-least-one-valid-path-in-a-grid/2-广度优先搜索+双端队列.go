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
	visited := make(map[[2]int]bool)
	deque := make([][]int, 0)
	deque = append(deque, []int{0, 0})
	for len(deque) > 0 {
		node := deque[0]
		deque = deque[1:]
		a, b := node[0], node[1]
		if visited[[2]int{a, b}] == true {
			continue
		}
		visited[[2]int{a, b}] = true
		for i := 0; i < 4; i++ {
			x, y := a+dx[i], b+dy[i]
			if 0 <= x && x < n && 0 <= y && y < m {
				dis := arr[a][b]
				if i+1 != grid[a][b] {
					dis = dis + 1
				}
				if dis < arr[x][y] {
					arr[x][y] = dis
					if i+1 == grid[a][b] { // 相同方向
						deque = append([][]int{{x, y}}, deque...) // 插入到前面
					} else {
						deque = append(deque, []int{x, y}) // 插入到后面
					}
				}
			}
		}
	}
	return arr[n-1][m-1]
}
