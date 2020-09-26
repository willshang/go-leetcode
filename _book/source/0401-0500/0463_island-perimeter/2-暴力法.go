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

// leetcode463_岛屿的周长
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4
			if i > 0 && grid[i-1][j] == 1 {
				res -= 2
			}
			if j > 0 && grid[i][j-1] == 1 {
				res -= 2
			}
		}
	}
	return res
}
