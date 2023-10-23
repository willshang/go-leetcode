package main

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{
		{1, 2, 2},
		{2, 3, 2},
	}, 0, 1, 3))
}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	dfs(grid, r0, c0, grid[r0][c0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				grid[i][j] = color // 染色
			}
		}
	}
	return grid
}

func dfs(grid [][]int, i, j int, color int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||
		grid[i][j] != color {
		return
	}
	grid[i][j] = -color // 先标记相反，表示访问过
	dfs(grid, i+1, j, color)
	dfs(grid, i-1, j, color)
	dfs(grid, i, j+1, color)
	dfs(grid, i, j-1, color)
	if 0 < i && i < len(grid)-1 && 0 < j && j < len(grid[0])-1 &&
		color == abs(grid[i-1][j]) && color == abs(grid[i+1][j]) &&
		color == abs(grid[i][j-1]) && color == abs(grid[i][j+1]) {
		grid[i][j] = color // 访问完，内部的联通分量还原，不做改变
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
