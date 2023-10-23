package main

import "fmt"

func main() {
	arr := [][]byte{
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
	}
	//arr := [][]byte{
	//	{'O', 'O', 'O'},
	//	{'O', 'O', 'O'},
	//	{'O', 'O', 'O'},
	//}
	//arr := [][]byte{
	//	{'X', 'X', 'X', 'X'},
	//	{'X', 'O', 'O', 'X'},
	//	{'X', 'X', 'O', 'X'},
	//	{'X', 'O', 'X', 'X'},
	//}
	//arr := [][]byte{
	//	{'X', 'O', 'X', 'O', 'X', 'O'},
	//	{'O', 'X', 'O', 'X', 'O', 'X'},
	//	{'X', 'O', 'X', 'O', 'X', 'O'},
	//	{'O', 'X', 'O', 'X', 'O', 'X'},
	//}
	solve(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(string(arr[i][j]))
		}
		fmt.Println()
	}
}

// leetcode130_被围绕的区域
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	n := len(board)
	m := len(board[0])
	fa = Init(n*m + 1)
	target := n * m
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == n-1 || j == 0 || j == m-1 {
					union(i*m+j, target)
				} else {
					if board[i-1][j] == 'O' {
						union(i*m+j, (i-1)*m+j)
					}
					if board[i+1][j] == 'O' {
						union(i*m+j, (i+1)*m+j)
					}
					if board[i][j-1] == 'O' {
						union(i*m+j, i*m+j-1)
					}
					if board[i][j+1] == 'O' {
						union(i*m+j, i*m+j+1)
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' && find(i*m+j) != find(target) {
				board[i][j] = 'X'
			}
		}
	}
}

var fa []int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// 查询
func find(x int) int {
	if fa[x] == x {
		return x
	}
	// 路径压缩
	fa[x] = find(fa[x])
	return fa[x]
}

// 合并
func union(i, j int) {
	fa[find(i)] = find(j)
}

func query(i, j int) bool {
	return find(i) == find(j)
}
