package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(arr)
	fmt.Println(arr)
}

// leetcode289_生命游戏
func gameOfLife(board [][]int) {
	// 0 00 => 死
	// 1 01 => 活
	// 2 10 => 死=>活
	// 3 11 => 活=>死
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			count := 0
			for a := i - 1; a <= i+1; a++ {
				for b := j - 1; b <= j+1; b++ {
					if a == i && b == j {
						continue
					}
					if 0 <= a && a < len(board) &&
						0 <= b && b < len(board[i]) {
						count = count + board[a][b]%2
					}
				}
			}
			if (count < 2 || count > 3) && board[i][j] == 1 {
				board[i][j] = 3
			}
			if count == 3 && board[i][j] == 0 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}
