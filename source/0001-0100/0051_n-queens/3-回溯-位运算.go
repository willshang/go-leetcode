package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	// 初始化棋盘
	arr := make([][]string, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, n)
		for j := 0; j < n; j++ {
			arr[i][j] = "."
		}
	}
	// 从第1行开始,上层是满足条件
	dfs(arr, 0, 0, 0, 0)
	return res
}

func dfs(arr [][]string, row int, rows, left, right int) {
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
		a := uint(col)
		b := uint(row - col + n - 1)
		c := uint(row + col)
		if ((rows>>a)&1) != 0 || ((left>>b)&1) != 0 || ((right>>c)&1) != 0 {
			continue
		}
		arr[row][col] = "Q"
		dfs(arr, row+1, rows^(1<<a), left^(1<<b), right^(1<<c))
		arr[row][col] = "."
	}
}
