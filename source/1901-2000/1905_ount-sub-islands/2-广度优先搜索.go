package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	arr1 := [][]int{
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
	}
	arr2 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}
	fmt.Println(countSubIslands(arr1, arr2))
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[i]); j++ {
			if grid2[i][j] == 1 { // 查找grid2
				if bfs(grid1, grid2, i, j) == true {
					res++
				}
			}
		}
	}
	return res
}

// leetcode200.岛屿数量
func bfs(grid1, grid2 [][]int, i, j int) bool {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{i, j})
	flag := true
	if grid1[i][j] == 0 {
		flag = false
	}
	grid1[i][j] = 0
	grid2[i][j] = 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := node[0] + dx[i]
			y := node[1] + dy[i]
			if 0 <= x && x < len(grid2) &&
				0 <= y && y < len(grid2[0]) &&
				grid2[x][y] == 1 {
				queue = append(queue, [2]int{x, y})
				if grid1[x][y] == 0 {
					flag = false
				}
				grid1[x][y] = 0
				grid2[x][y] = 0
			}
		}
	}
	return flag
}
