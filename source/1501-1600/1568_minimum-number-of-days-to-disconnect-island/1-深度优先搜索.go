package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{0, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0},
	}
	fmt.Println(minDays(arr))
}

// leetcode1568_使陆地分离的最少天数
func minDays(grid [][]int) int {
	temp := copyArr(grid)
	nums := numIslands(temp)
	if nums >= 2 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			temp = copyArr(grid)
			temp[i][j] = 0
			if numIslands(temp) == 2 {
				return 1
			}
		}
	}
	return 2
}

func numIslands(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) ||
		grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func copyArr(grid [][]int) [][]int {
	temp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		temp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			temp[i][j] = grid[i][j]
		}
	}
	return temp
}
