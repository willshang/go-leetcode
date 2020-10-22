package main

import "fmt"

func main() {
	nums := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(nums))
}

func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return 0
}

func dfs(grid [][]int, i, j int) int {
	// 边界+1
	if !(0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])) {
		return 1
	}
	// 水域+1
	if grid[i][j] == 0 {
		return 1
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 2
	return dfs(grid, i-1, j) +
		dfs(grid, i+1, j) +
		dfs(grid, i, j-1) +
		dfs(grid, i, j+1)
}
