package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(arr)
	fmt.Println(arr)
}

// leetcode73_矩阵置零
func setZeroes(matrix [][]int) {
	flag := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 第一列处理
	if flag == true {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
