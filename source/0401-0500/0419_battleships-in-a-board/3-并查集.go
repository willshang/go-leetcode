package main

import "fmt"

func main() {
	arr := [][]byte{
		{'X', '.', 'X'},
		{'X', '.', 'X'},
	}
	fmt.Println(countBattleships(arr))
}

func countBattleships(board [][]byte) int {
	a, b := len(board), len(board[0])
	fa = Init(a * b)
	m := make(map[int]bool)
	arr := make([]int, 0)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if board[i][j] == 'X' {
				if i < a-1 && board[i][j] == board[i+1][j] {
					union(i*b+j, i*b+b+j)
				}
				if j < b-1 && board[i][j] == board[i][j+1] {
					union(i*b+j, i*b+j+1)
				}
				arr = append(arr, i*b+j)
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		m[find(arr[i])] = true
	}
	return len(m)
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

func union(a, b int) {
	x, y := find(a), find(b)
	if x != y {
		fa[x] = y
	}
}

// 彻底路径压缩
func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}
