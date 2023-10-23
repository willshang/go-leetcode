package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 3},
		{0, -4},
	}
	fmt.Println(maxProductPath(arr))
}

// leetcode1594_矩阵的最大非负积
var res int

func maxProductPath(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res = -1
	dfs(grid, 0, 0, 1)
	if res < 0 {
		return -1
	}
	return res % 1000000007
}

func dfs(grid [][]int, i, j int, value int) {
	value = value * grid[i][j]
	if grid[i][j] == 0 {
		if value > res {
			res = value
		}
		return
	}
	if i == len(grid)-1 && j == len(grid[0])-1 {
		if value > res {
			res = value
		}
	}
	if i < len(grid)-1 {
		dfs(grid, i+1, j, value)
	}
	if j < len(grid[0])-1 {
		dfs(grid, i, j+1, value)
	}
}
