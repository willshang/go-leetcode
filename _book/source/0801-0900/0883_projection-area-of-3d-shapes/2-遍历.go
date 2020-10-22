package main

import "fmt"

func main() {
	area := [][]int{{1, 2}, {3, 4}}
	fmt.Println(projectionArea(area))
}

func projectionArea(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		yz := 0
		xz := 0
		// 每一行最大值之和，每一列最大值之和
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 0 {
				res++
			}
			if yz < grid[i][j] {
				yz = grid[i][j]
			}
			if xz < grid[j][i] {
				xz = grid[j][i]
			}
		}
		res = res + yz + xz
	}
	return res
}
