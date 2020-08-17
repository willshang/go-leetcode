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

// leetcode37_解数独
var rows, cols, arrs [9][9]int

func solveSudoku(board [][]byte) {
	rows = [9][9]int{}
	cols = [9][9]int{}
	arrs = [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				index := (i/3)*3 + j/3
				rows[i][num] = 1
				cols[j][num] = 1
				arrs[index][num] = 1
			}
		}
	}
	dfs(board, 0)
}
func dfs(board [][]byte, index int) bool {
	if index == 81 {
		return true
	}
	row := index / 9
	col := index % 9
	c := (row/3)*3 + col/3
	if board[row][col] != '.' {
		return dfs(board, index+1)
	}
	for i := 0; i < 9; i++ {
		if rows[row][i] == 1 || cols[col][i] == 1 || arrs[c][i] == 1 {
			continue
		}
		board[row][col] = byte(i + '1')
		rows[row][i], cols[col][i], arrs[c][i] = 1, 1, 1
		if dfs(board, index+1) == true {
			return true
		}
		rows[row][i], cols[col][i], arrs[c][i] = 0, 0, 0
		board[row][col] = '.'
	}
	return false
}
