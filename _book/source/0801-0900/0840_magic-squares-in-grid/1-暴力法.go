package main

import "fmt"

func main() {
	arr := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(arr))
}

// leetcode840_矩阵中的幻方
func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			if grid[i+1][j+1] != 5 {
				continue
			}
			if !available(i, j, grid) {
				continue
			}
			if grid[i][j]+grid[i][j+1]+grid[i][j+2] == 15 &&
				grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] == 15 &&
				grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] == 15 &&
				grid[i][j]+grid[i+1][j]+grid[i+2][j] == 15 &&
				grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] == 15 &&
				grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] == 15 &&
				grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] == 15 &&
				grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] == 15 {
				res++
			}
		}
	}
	return res
}

func available(x, y int, g [][]int) bool {
	tmp := [16]int{}
	for i := x; i <= x+2; i++ {
		for j := y; j <= y+2; j++ {
			tmp[g[i][j]]++
		}
	}

	for i := 1; i <= 9; i++ {
		if tmp[i] != 1 {
			return false
		}
	}
	return true
}
