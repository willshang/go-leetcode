package main

func main() {

}

// leetcode794_有效的井字游戏
func validTicTacToe(board []string) bool {
	var XCount, OCount int
	n := len(board)
	rows := make([][2]int, n)         // 行
	cols := make([][2]int, n)         // 列
	left, right := [2]int{}, [2]int{} // 对角线
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				XCount++
				rows[i][0]++
				cols[j][0]++
				if i == j {
					left[0]++
				}
				if i == n-1-j {
					right[0]++
				}
			} else if board[i][j] == 'O' {
				OCount++
				rows[i][1]++
				cols[j][1]++
				if i == j {
					left[1]++
				}
				if i == n-1-j {
					right[1]++
				}
			}
		}
	}
	if XCount != OCount && XCount-1 != OCount {
		return false
	}
	for i := 0; i < n; i++ { // 行列判断
		if (rows[i][0] == n || cols[i][0] == n) && XCount-1 != OCount {
			return false
		}
		if (rows[i][1] == n || cols[i][1] == n) && XCount != OCount {
			return false
		}
	}
	if (left[0] == n || right[0] == n) && XCount-1 != OCount { // 对角线判断
		return false
	}
	if (left[1] == n || right[1] == n) && XCount != OCount {
		return false
	}
	return true
}
