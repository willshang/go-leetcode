package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(updateMatrix(arr))
}

// 剑指OfferII107.矩阵中的距离
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	queue := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := node[0] + dx[i]
			y := node[1] + dy[i]
			if 0 <= x && x < n && 0 <= y && y < m && matrix[x][y] == -1 {
				matrix[x][y] = matrix[node[0]][node[1]] + 1
				queue = append(queue, [2]int{x, y})
			}
		}
	}
	return matrix
}
