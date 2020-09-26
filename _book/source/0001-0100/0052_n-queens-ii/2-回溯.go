package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
}

var res int

func totalNQueens(n int) int {
	res = 0
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
	if len(arr) == row {
		res++
		return
	}
	// 每列尝试
	for col := 0; col < len(arr[0]); col++ {
		if valid(arr, row, col) == false {
			continue
		}
		arr[row][col] = "Q"
		dfs(arr, row+1)
		arr[row][col] = "."
	}
}

func valid(arr [][]string, row, col int) bool {
	n := len(arr)
	// 当前列判断(竖着)
	for row := 0; row < n; row++ {
		if arr[row][col] == "Q" {
			return false
		}
	}
	// 左上角
	for row, col := row-1, col-1; row >= 0 && col >= 0; row, col = row-1, col-1 {
		if arr[row][col] == "Q" {
			return false
		}
	}
	// 右上角
	for row, col := row-1, col+1; row >= 0 && col < n; row, col = row-1, col+1 {
		if arr[row][col] == "Q" {
			return false
		}
	}
	return true
}
