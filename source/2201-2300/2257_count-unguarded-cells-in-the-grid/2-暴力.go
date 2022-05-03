package main

import "fmt"

func main() {
	arr1 := [][]int{
		{0, 0},
		{1, 1},
		{2, 3},
	}
	arr2 := [][]int{
		{0, 1},
		{2, 2},
		{1, 4},
	}
	fmt.Println(countUnguarded(4, 6, arr1, arr2))
}

// leetcode2257_统计网格图中没有被保卫的格子数
// 顺时针：上右下左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < len(walls); i++ {
		a, b := walls[i][0], walls[i][1]
		arr[a][b] = 1
	}
	for i := 0; i < len(guards); i++ {
		a, b := guards[i][0], guards[i][1]
		arr[a][b] = 2
	}
	for i := 0; i < len(guards); i++ {
		a, b := guards[i][0], guards[i][1]
		for k := 0; k < 4; k++ {
			x, y := a+dx[k], b+dy[k]
			for 0 <= x && x < m && 0 <= y && y < n && arr[x][y] != 1 && arr[x][y] != 2 {
				arr[x][y] = 3
				x, y = x+dx[k], y+dy[k]
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				res++
			}
		}
	}
	return res
}
