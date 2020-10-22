package main

func main() {

}

var dx = []int{-1, 1, 0, 0, 1, 1, -1, -1}
var dy = []int{0, 0, -1, 1, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		bfs(board, x, y)
	}
	return board
}

func bfs(board [][]byte, x, y int) {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{x, y})
	visited[x][y] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count := 0
		a := node[0]
		b := node[1]
		for j := 0; j < 8; j++ {
			newX := dx[j] + a
			newY := dy[j] + b
			if newX < 0 || newX >= len(board) ||
				newY < 0 || newY >= len(board[0]) ||
				visited[newX][newY] == true {
				continue
			}
			if board[newX][newY] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[a][b] = byte(count + '0')
		} else {
			board[a][b] = 'B'
			for j := 0; j < 8; j++ {
				newX := dx[j] + a
				newY := dy[j] + b
				if newX < 0 || newX >= len(board) ||
					newY < 0 || newY >= len(board[0]) ||
					board[newX][newY] != 'E' ||
					visited[newX][newY] == true {
					continue
				}
				queue = append(queue, [2]int{newX, newY})
				visited[newX][newY] = true
			}
		}
	}
}
