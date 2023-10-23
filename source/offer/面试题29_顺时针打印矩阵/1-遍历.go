package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(arr))
}

var res []int

func spiralOrder(matrix [][]int) []int {
	res = make([]int, 0)
	rows := len(matrix)
	if rows == 0 {
		return res
	}
	cols := len(matrix[0])
	if cols == 0 {
		return res
	}
	start := 0
	for cols > start*2 && rows > start*2 {
		printCircle(matrix, cols, rows, start)
		start++
	}
	return res
}

func printCircle(matrix [][]int, cols, rows, start int) {
	x := cols - 1 - start
	y := rows - 1 - start
	// 左到右
	for i := start; i <= x; i++ {
		res = append(res, matrix[start][i])
	}
	// 上到下
	if start < y {
		for i := start + 1; i <= y; i++ {
			res = append(res, matrix[i][x])
		}
	}
	// 右到左
	if start < x && start < y {
		for i := x - 1; i >= start; i-- {
			res = append(res, matrix[y][i])
		}
	}
	// 下到上
	if start < x && start < y-1 {
		for i := y - 1; i >= start+1; i-- {
			res = append(res, matrix[i][start])
		}
	}
}
