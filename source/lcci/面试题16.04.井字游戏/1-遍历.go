package main

func main() {

}

// 程序员面试金典16.04_井字游戏
func tictactoe(board []string) string {
	n := len(board)
	flag := false                     // 有没有空格
	rows := make([][2]int, n)         // 行
	cols := make([][2]int, n)         // 列
	left, right := [2]int{}, [2]int{} // 对角线
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == ' ' {
				flag = true
			} else if board[i][j] == 'X' {
				rows[i][0]++
				cols[j][0]++
				if i == j {
					left[0]++
				}
				if i == n-1-j {
					right[0]++
				}
			} else if board[i][j] == 'O' {
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
	for i := 0; i < n; i++ { // 行列判断
		if rows[i][0] == n || cols[i][0] == n {
			return "X"
		}
		if rows[i][1] == n || cols[i][1] == n {
			return "O"
		}
	}
	if left[0] == n || right[0] == n { // 对角线判断
		return "X"
	}
	if left[1] == n || right[1] == n {
		return "O"
	}
	if flag == true {
		return "Pending"
	}
	return "Draw"
}
