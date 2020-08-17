package main

import "fmt"

func main() {
	arr := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(arr)
	fmt.Println(arr)
}

// leetcode130_被围绕的区域
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1) &&
				board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) ||
		board[i][j] == '#' || board[i][j] == 'X' {
		return
	}
	board[i][j] = '#'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}
