package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	fmt.Println(numEnclaves(arr))
}

// leetcode1020_飞地的数量
func numEnclaves(A [][]int) int {
	for i := 0; i < len(A); i++ {
		dfs(A, i, 0)
		dfs(A, i, len(A[i])-1)
	}
	for i := 0; i < len(A[0]); i++ {
		dfs(A, 0, i)
		dfs(A, len(A)-1, i)
	}
	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs(A [][]int, i, j int) {
	if i < 0 || i >= len(A) || j < 0 || j >= len(A[i]) {
		return
	}
	if A[i][j] == 0 {
		return
	}
	A[i][j] = 0
	dfs(A, i, j+1)
	dfs(A, i, j-1)
	dfs(A, i+1, j)
	dfs(A, i-1, j)
}
