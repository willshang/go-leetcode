package main

func main() {

}

// leetcode_lcp41_黑白翻转棋
var res int
var n, m int
var dx = []int{-1, 1, 0, 0, 1, 1, -1, -1}
var dy = []int{0, 0, -1, 1, 1, -1, -1, 1}

func flipChess(chessboard []string) int {
	res = 0
	n, m = len(chessboard), len(chessboard[0])
	temp := make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if chessboard[i][j] == '.' {
				for a := 0; a < n; a++ {
					temp[a] = []byte(chessboard[a])
				}
				temp[i][j] = 'X'
				dfs(temp, i, j) // 尝试在i,j位置下黑棋
				count := 0
				for a := 0; a < n; a++ { // 统计结果
					for b := 0; b < m; b++ {
						if temp[a][b] == 'X' && chessboard[a][b] == 'O' {
							count++
						}
					}
				}
				res = max(res, count) // 更新结果
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(arr [][]byte, i, j int) {
	arr[i][j] = 'X'
	for k := 0; k < 8; k++ {
		if ok, path := judge(arr, i, j, 'X', dx[k], dy[k]); ok == true {
			for c := 0; c < len(path); c++ {
				a, b := path[c][0], path[c][1]
				arr[a][b] = 'X'
				dfs(arr, a, b)
			}
		}
	}
}

// leetcode1958.检查操作是否合法
func judge(board [][]byte, rMove int, cMove int, color byte, dirX, dirY int) (res bool, path [][]int) {
	x, y := rMove+dirX, cMove+dirY
	count := 1
	for 0 <= x && x < n && 0 <= y && y < m {
		if board[x][y] == '.' {
			return false, nil
		}
		path = append(path, []int{x, y})
		if count == 1 {
			if board[x][y] == color {
				return false, nil
			}
		} else {
			if board[x][y] == color {
				return true, path
			}
		}
		count++
		x = x + dirX
		y = y + dirY
	}
	return false, nil
}
