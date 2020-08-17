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

func countNegatives(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := len(grid[i]) - 1; j >= -1; j-- {
			if j == -1 {
				res = res + len(grid[i])
				break
			}
			if grid[i][j] >= 0 {
				count := len(grid[i]) - 1 - j
				res = res + count
				break
			}
		}
	}
	return res
}
