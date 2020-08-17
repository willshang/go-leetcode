package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
}

// leetcode52_N皇后II
var res int

func totalNQueens(n int) int {
	res = 0
	// 从第1行开始,上层是满足条件
	dfs(0, n, 0, 0, 0)
	return res
}

func dfs(row, n int, rows, left, right int) {
	if n == row {
		res++
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
		dfs(row+1, n, rows^(1<<a), left^(1<<b), right^(1<<c))
	}
}
