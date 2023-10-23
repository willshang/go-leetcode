package main

func main() {

}

// leetcode1219_黄金矿工
var res int

func getMaximumGold(grid [][]int) int {
	res = 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				dfs(grid, i, j, 0, visited)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int, sum int, visited [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||
		grid[i][j] == 0 || visited[i][j] == true {
		return
	}
	visited[i][j] = true
	sum = sum + grid[i][j]
	if sum > res {
		res = sum
	}
	dfs(grid, i+1, j, sum, visited)
	dfs(grid, i-1, j, sum, visited)
	dfs(grid, i, j+1, sum, visited)
	dfs(grid, i, j-1, sum, visited)
	visited[i][j] = false
}
