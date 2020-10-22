package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println(pondSizes(arr))
}

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
	if grid[i][j] != 0 {
		return 0
	}
	grid[i][j] = 1
	area := 1
	for a := i - 1; a <= i+1; a++ {
		for b := j - 1; b <= j+1; b++ {
			if (i == a && j == b) || a < 0 || a >= len(grid) ||
				b < 0 || b >= len(grid[0]) {
				continue
			}
			area = area + getArea(grid, a, b)
		}
	}
	return area
}
