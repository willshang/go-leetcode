package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(arr)
	fmt.Println(arr)
}

// 程序员面试金典01.07_旋转矩阵
func rotate(matrix [][]int) {
	n := len(matrix)
	// 同行逆置
	// [[1 2 3] [4 5 6] [7 8 9]]
	// [[3 2 1] [6 5 4] [9 8 7]]
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	// 左下右上对角线对互换
	// [[3 2 1] [6 5 4] [9 8 7]]
	// [[7 4 1] [8 5 2] [9 6 3]]
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
}
