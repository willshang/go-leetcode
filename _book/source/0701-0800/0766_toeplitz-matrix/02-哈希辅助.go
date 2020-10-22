package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	fmt.Println(isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	m := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if value, ok := m[i-j]; ok {
				if matrix[i][j] != value {
					return false
				}
			} else {
				m[i-j] = matrix[i][j]
			}
		}
	}
	return true
}
