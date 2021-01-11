package main

func main() {

}

// leetcode419_甲板上的战舰
func countBattleships(board [][]byte) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				if (i > 0 && board[i-1][j] == 'X') ||
					(j > 0 && board[i][j-1] == 'X') {
					continue
				}
				res++
			}
		}
	}
	return res
}
