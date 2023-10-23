package main

func main() {

}

// leetcode980_不同路径III
var res int
var n, m int

func uniquePathsIII(grid [][]int) int {
	res = 0
	n, m = len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	x, y := 0, 0
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				x, y = i, j
				continue
			}
			if grid[i][j] == 0 {
				count++
			}
		}
	}
	dfs(grid, x, y, visited, count)
	return res
}

func dfs(grid [][]int, i, j int, visited [][]bool, count int) {
	if i < 0 || i >= n || j < 0 || j >= m ||
		grid[i][j] == -1 || visited[i][j] == true {
		return
	}
	if grid[i][j] == 2 {
		if count == -1 { // 包括起始点1
			res++
		}
		return
	}
	visited[i][j] = true
	dfs(grid, i, j+1, visited, count-1)
	dfs(grid, i, j-1, visited, count-1)
	dfs(grid, i+1, j, visited, count-1)
	dfs(grid, i-1, j, visited, count-1)
	visited[i][j] = false
}
