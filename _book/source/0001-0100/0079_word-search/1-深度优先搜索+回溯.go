package main

import "fmt"

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	fmt.Println(exist(board, "abcd"))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, level int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		board[i][j] != word[level] {
		return false
	}
	if level == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = ' '
	res := dfs(board, i+1, j, word, level+1) ||
		dfs(board, i-1, j, word, level+1) ||
		dfs(board, i, j+1, word, level+1) ||
		dfs(board, i, j-1, word, level+1)
	board[i][j] = temp
	return res
}
