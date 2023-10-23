package main

import "fmt"

func main() {
	arr := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'B', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', 'W', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'W', 'B', '.', '.', '.'},
		{'.', '.', '.', '.', 'W', '.', '.', '.'},
		{'.', '.', '.', '.', 'B', 'W', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'W', '.'},
		{'.', '.', '.', '.', '.', '.', '.', 'B'},
	}

	fmt.Println(checkMove(arr, 4, 4, 'W'))
}

// leetcode1958_检查操作是否合法
var dx = []int{1, 1, 0, -1, -1, -1, 0, 1}
var dy = []int{0, 1, 1, 1, 0, -1, -1, -1}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	for i := 0; i < 8; i++ {
		if judge(board, rMove, cMove, color, dx[i], dy[i]) == true {
			return true
		}
	}
	return false
}

func judge(board [][]byte, rMove int, cMove int, color byte, dirX, dirY int) bool {
	x, y := rMove+dirX, cMove+dirY
	count := 1
	for 0 <= x && x < 8 && 0 <= y && y < 8 {
		if board[x][y] == '.' {
			return false
		}
		if count == 1 {
			if board[x][y] == color {
				return false
			}
		} else {
			if board[x][y] == color {
				return true
			}
		}
		count++
		x = x + dirX
		y = y + dirY
	}
	return false
}
