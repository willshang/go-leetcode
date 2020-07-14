package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(arr))
}

// leetcode54_螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	rows := len(matrix)
	if rows == 0 {
		return res
	}
	cols := len(matrix[0])
	if cols == 0 {
		return res
	}
	x1, x2, y1, y2 := 0, rows-1, 0, cols-1
	direct := 0
	for x1 <= x2 && y1 <= y2 {
		direct = (direct + 4) % 4
		if direct == 0 {
			for i := y1; i <= y2; i++ {
				res = append(res, matrix[x1][i])
			}
			x1++
		} else if direct == 1 {
			for i := x1; i <= x2; i++ {
				res = append(res, matrix[i][y2])
			}
			y2--
		} else if direct == 2 {
			for i := y2; i >= y1; i-- {
				res = append(res, matrix[x2][i])
			}
			x2--
		} else if direct == 3 {
			for i := x2; i >= x1; i-- {
				res = append(res, matrix[i][y1])
			}
			y1++
		}
		direct++
	}
	return res
}
