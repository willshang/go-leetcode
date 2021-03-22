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
	i := 0
	j := len(grid[0]) - 1
	for i < len(grid) && j >= 0 {
		if grid[i][j] >= 0 {
			res = res + len(grid[0]) - j - 1
			i++
		} else {
			j--
		}
	}
	if j < 0 {
		res = res + (len(grid)-i)*len(grid[0])
	}
	return res
}
