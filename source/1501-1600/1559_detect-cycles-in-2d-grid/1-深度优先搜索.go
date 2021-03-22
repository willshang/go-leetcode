package main

import "fmt"

func main() {
	arr := [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'a', 'a', 'a'},
	}
	fmt.Println(containsCycle(arr))
}

// leetcode1559_二维网格图中探测环
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

var visited map[[2]int]bool

func containsCycle(grid [][]byte) bool {
	n, m := len(grid), len(grid[0])
	visited = make(map[[2]int]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[[2]int{i, j}] == true {
				continue
			}
			start := grid[i][j]
			if dfs(grid, start, i, j, -1, -1) == true {
				return true
			}
		}
	}
	return false
}

// x,y 当前节点， pX,pY 父节点
func dfs(grid [][]byte, start byte, x, y int, pX, pY int) bool {
	visited[[2]int{x, y}] = true
	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) ||
			(newX == pX && newY == pY) {
			continue
		}
		if start == grid[newX][newY] {
			if visited[[2]int{newX, newY}] == true {
				return true
			}
			result := dfs(grid, start, newX, newY, x, y)
			if result == true {
				return true
			}
		}
	}
	return false
}
