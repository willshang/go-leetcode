package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(arr))
}

var arr [][]int

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	arr = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	return dfs(grid, n-1, m-1)
}

func dfs(grid [][]int, n, m int) int {
	if m == 0 && n == 0 {
		arr[0][0] = grid[0][0]
		return grid[0][0]
	}
	if n == 0 {
		return grid[0][m] + dfs(grid, 0, m-1)
	}
	if m == 0 {
		return grid[n][0] + dfs(grid, n-1, 0)
	}
	if arr[n][m] > 0 {
		return arr[n][m]
	}
	arr[n][m] = min(dfs(grid, n-1, m), dfs(grid, n, m-1)) + grid[n][m]
	return arr[n][m]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
