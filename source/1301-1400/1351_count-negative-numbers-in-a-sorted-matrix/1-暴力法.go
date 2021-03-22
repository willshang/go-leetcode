package main

import "fmt"

func main() {
	arr := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(arr))
}

// leetcode1351_统计有序矩阵中的负数
func countNegatives(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				res++
			}
		}
	}
	return res
}
