package main

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{
		{1, 2, 2},
		{2, 3, 2},
	}, 0, 1, 3))
}

// leetcode1034_边框着色
var visited []bool

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	visited = make([]bool, len(grid)*len(grid[0]))
	dfs(grid, r0, c0, grid[r0][c0], color)
	return grid
}

func dfs(grid [][]int, i, j int, targetColor int, color int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) { // 边界
		return 0
	}
	if visited[i*len(grid[0])+j] == true { // 先判断，因为修改了颜色
		return 1
	}
	if grid[i][j] != targetColor {
		return 0
	}
	visited[i*len(grid[0])+j] = true
	a := dfs(grid, i+1, j, targetColor, color)
	b := dfs(grid, i-1, j, targetColor, color)
	c := dfs(grid, i, j+1, targetColor, color)
	d := dfs(grid, i, j-1, targetColor, color)
	if a+b+c+d < 4 {
		grid[i][j] = color
	}
	return 1
}
