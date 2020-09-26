package main

import "fmt"

func main() {
	arr := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(string(arr[i][j]))
		}
		fmt.Println()
	}
	fmt.Println(arr)
}

func solveSudoku(board [][]byte) {
	dfs(board, 0)
}

func dfs(board [][]byte, index int) bool {
	if index == 81 {
		return true
	}
	row := index / 9
	col := index % 9
	if board[row][col] != '.' {
		return dfs(board, index+1)
	}
	for i := 0; i < 9; i++ {
		board[row][col] = byte(i + '1')
		if isValidSudoku(board) == false {
			board[row][col] = '.'
			continue
		}
		if dfs(board, index+1) == true {
			return true
		}
		board[row][col] = '.'
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	var row, col, arr [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				index := (i/3)*3 + j/3
				if row[i][num] == 1 || col[j][num] == 1 || arr[index][num] == 1 {
					return false
				}
				row[i][num] = 1
				col[j][num] = 1
				arr[index][num] = 1
			}
		}
	}
	return true
}
