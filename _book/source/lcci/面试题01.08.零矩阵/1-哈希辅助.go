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
	x := make(map[int]int)
	y := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				x[i] = 1
				y[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if x[i] == 1 || y[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
