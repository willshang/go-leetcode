package main

import "fmt"

func main() {
	arr := [][]int{{0, 1, 0, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(arr))
}

// 剑指OfferII105.岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				value := dfs(grid, i, j)
				if value > res {
					res = value
				}
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) ||
		grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	res := 1
	res = res + dfs(grid, i+1, j)
	res = res + dfs(grid, i-1, j)
	res = res + dfs(grid, i, j+1)
	res = res + dfs(grid, i, j-1)
	return res
}
