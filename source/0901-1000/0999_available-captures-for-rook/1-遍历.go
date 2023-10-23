package main

import "fmt"

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(numRookCaptures(board))
}

func numRookCaptures(board [][]byte) int {
	res := 0
	var x, y int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				x = i
				y = j
				break
			}
		}
	}
	// 向右
	for i := y; i < 8 && board[x][i] != 'B'; i++ {
		if board[x][i] == 'p' {
			res++
			break
		}
	}
	// 向左
	for i := y; i >= 0 && board[x][i] != 'B'; i-- {
		if board[x][i] == 'p' {
			res++
			break
		}
	}
	// 向下
	for i := x; i < 8 && board[i][y] != 'B'; i++ {
		if board[i][y] == 'p' {
			res++
			break
		}
	}
	// 向上
	for i := x; i >= 0 && board[i][y] != 'B'; i-- {
		if board[i][y] == 'p' {
			res++
			break
		}
	}
	return res
}
