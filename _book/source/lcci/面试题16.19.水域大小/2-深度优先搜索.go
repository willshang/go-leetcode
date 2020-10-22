package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println(pondSizes(arr))
}

// 程序员面试金典16.19_水域大小
func pondSizes(land [][]int) []int {
	res := make([]int, 0)
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 0 {
				res = append(res, getArea(land, i, j))
			}
		}
	}
	sort.Ints(res)
	return res
}

func getArea(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) ||
		j < 0 || j >= len(grid[0]) || grid[i][j] != 0 {
		return 0
	}

	grid[i][j] = 1
	res := 1
	res = res + getArea(grid, i+1, j)
	res = res + getArea(grid, i+1, j+1)
	res = res + getArea(grid, i+1, j-1)
	res = res + getArea(grid, i-1, j)
	res = res + getArea(grid, i-1, j+1)
	res = res + getArea(grid, i-1, j-1)
	res = res + getArea(grid, i, j+1)
	res = res + getArea(grid, i, j-1)
	return res
}
