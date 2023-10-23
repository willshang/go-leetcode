package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}

// leetcode885_螺旋矩阵III
// 顺时针：上右下左
// 本题：右、下、左、上
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	res := make([][]int, 0)
	total := rows * cols
	x, y := rStart, cStart
	index := 0
	step := 2
	dir := 0
	for index < total {
		for i := 0; i < step/2; i++ { // 本次移动的步数
			if 0 <= x && x < rows && 0 <= y && y < cols {
				res = append(res, []int{x, y})
				index++
				if index == total {
					return res
				}
			}
			x = x + dx[dir]
			y = y + dy[dir]
		}
		dir = (dir + 1) % 4
		step++
	}
	return res
}
