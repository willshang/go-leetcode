package main

import "fmt"

func main() {
	//arr := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	arr := [][]int{
		{1},
		{2},
		{3},
		{4},
	}
	fmt.Println(shiftGrid(arr, 1))
}

// leetcode1260_二维网格迁移
func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := ((i*m + j) + k) % (n * m) / m
			y := ((i*m + j) + k) % (n * m) % m
			// x := (i + (j+k)/m) % n
			// y := (j + k) % m
			res[x][y] = grid[i][j]
		}
	}
	return res
}
