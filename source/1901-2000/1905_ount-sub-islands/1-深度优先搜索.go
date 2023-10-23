package main

import "fmt"

func main() {
	arr1 := [][]int{
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
	}
	arr2 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}
	fmt.Println(countSubIslands(arr1, arr2))
}

// leetcode1905_统计子岛屿
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[i]); j++ {
			if grid2[i][j] == 1 { // 查找grid2
				if dfs(grid1, grid2, i, j) == true {
					res++
				}
			}
		}
	}
	return res
}

// leetcode200.岛屿数量
func dfs(grid1, grid2 [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) || grid2[i][j] == 0 {
		return true
	}
	if grid1[i][j] == 0 {
		return false
	}
	grid1[i][j], grid2[i][j] = 0, 0
	res1 := dfs(grid1, grid2, i+1, j)
	res2 := dfs(grid1, grid2, i-1, j)
	res3 := dfs(grid1, grid2, i, j+1)
	res4 := dfs(grid1, grid2, i, j-1)
	if res1 == false || res2 == false || res3 == false || res4 == false {
		return false
	}
	return true
}
