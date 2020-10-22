package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

// leetcode51_N皇后
var res [][]string
var rows, left, right []bool

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
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
		temp := make([]string, 0)
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < n; j++ {
				str = str + arr[i][j]
			}
			temp = append(temp, str)
		}
		res = append(res, temp)
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
