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
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 第一行处理
	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	// 第一列处理
	if flag == true {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
