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

// leetcode999_可以被一步捕获的棋子数
func numRookCaptures(board [][]byte) int {
	res := 0
	var x, y int
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				x = i
				y = j
				break
			}
		}
	}
	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		for newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) {
			if board[newX][newY] == 'B' {
				break
			}
			if board[newX][newY] == 'p' {
				res++
				break
			}
			newX = newX + dx[i]
			newY = newY + dy[i]
		}
	}
	return res
}
