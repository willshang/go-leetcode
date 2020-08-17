package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
}

var res int
var rows, left, right []bool

func totalNQueens(n int) int {
	res = 0
	rows, left, right = make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	// 初始化棋盘
	arr := make([][]string, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, n)
		for j := 0; j < n; j++ {
			arr[i][j] = "."
		}
	}
	// 从第1行开始,上层是满足条件
	dfs(arr, 0)
	return res
}

func dfs(arr [][]string, row int) {
	n := len(arr)
	if len(arr) == row {
		res++
		return
	}
	// 每列尝试
	for col := 0; col < n; col++ {
		if rows[col] == true || left[row-col+n-1] == true || right[row+col] == true {
			continue
		}
		rows[col], left[row-col+n-1], right[row+col] = true, true, true
		arr[row][col] = "Q"
		dfs(arr, row+1)
		arr[row][col] = "."
		rows[col], left[row-col+n-1], right[row+col] = false, false, false
	}
}
