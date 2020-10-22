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

func gameOfLife(board [][]int) {
	temp := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		temp[i] = make([]int, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			count := 0
			for a := i - 1; a <= i+1; a++ {
				for b := j - 1; b <= j+1; b++ {
					if a == i && b == j {
						continue
					}
					if 0 <= a && a < len(board) &&
						0 <= b && b < len(board[i]) && board[a][b] == 1 {
						count++
					}
				}
			}
			temp[i][j] = board[i][j]
			if count < 2 || count > 3 {
				temp[i][j] = 0
			}
			if count == 3 && board[i][j] == 0 {
				temp[i][j] = 1
			}
		}
	}
	copy(board, temp)
}
