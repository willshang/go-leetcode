package main

func main() {

}

func countBattleships(board [][]byte) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				res++
				board[i][j] = '.'
				for x := i + 1; x < len(board); x++ {
					if board[x][j] == 'X' {
						board[x][j] = '.'
					} else {
						break
					}
				}
				for y := j + 1; y < len(board[i]); y++ {
					if board[i][y] == 'X' {
						board[i][y] = '.'
					} else {
						break
					}
				}
			}
		}
	}
	return res
}
