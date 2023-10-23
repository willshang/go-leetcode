package main

func main() {

}

// leetcode529_扫雷游戏
var dx = []int{-1, 1, 0, 0, 1, 1, -1, -1}
var dy = []int{0, 0, -1, 1, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(board, x, y)
	}
	return board
}

func dfs(board [][]byte, x, y int) {
	count := 0
	for i := 0; i < 8; i++ {
		newX := dx[i] + x
		newY := dy[i] + y
		if newX < 0 || newX >= len(board) || newY < 0 || newY >= len(board[0]) {
			continue
		}
		if board[newX][newY] == 'M' {
			count++
		}
	}
	if count > 0 {
		board[x][y] = byte(count + '0')
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			newX := dx[i] + x
			newY := dy[i] + y
			if newX < 0 || newX >= len(board) ||
				newY < 0 || newY >= len(board[0]) ||
				board[newX][newY] != 'E' {
				continue
			}
			dfs(board, newX, newY)
		}
	}
}
