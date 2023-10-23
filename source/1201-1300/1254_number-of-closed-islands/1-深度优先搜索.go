package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	fmt.Println(closedIsland(arr))
}

// leetcode1254_统计封闭岛屿的数目
func closedIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 && dfs(grid, i, j) == true {
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}
	if grid[i][j] == 1 {
		return true
	}
	grid[i][j] = 1
	up := dfs(grid, i, j+1)
	down := dfs(grid, i, j-1)
	left := dfs(grid, i-1, j)
	right := dfs(grid, i+1, j)
	return up && down && left && right
}
