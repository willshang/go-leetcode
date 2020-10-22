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

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func numEnclaves(A [][]int) int {
	queue := make([][2]int, 0)
	for i := 0; i < len(A); i++ {
		if A[i][0] == 1 {
			A[i][0] = 0
			queue = append(queue, [2]int{i, 0})
		}
		if A[i][len(A[i])-1] == 1 {
			A[i][len(A[i])-1] = 0
			queue = append(queue, [2]int{i, len(A[i]) - 1})
		}
	}
	for i := 0; i < len(A[0]); i++ {
		if A[0][i] == 1 {
			A[0][i] = 0
			queue = append(queue, [2]int{0, i})
		}
		if A[len(A)-1][i] == 1 {
			A[len(A)-1][i] = 0
			queue = append(queue, [2]int{len(A) - 1, i})
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			x := dx[k] + node[0]
			y := dy[k] + node[1]
			if x < 0 || x >= len(A) || y < 0 || y >= len(A[0]) {
				continue
			}
			if A[x][y] == 0 {
				continue
			}
			queue = append(queue, [2]int{x, y})
			A[x][y] = 0
		}
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
