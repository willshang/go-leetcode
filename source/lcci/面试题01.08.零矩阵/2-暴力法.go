package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(arr)
	fmt.Println(arr)
}

func setZeroes(matrix [][]int) {
	m := make(map[[2]int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MinInt32 {
				m[[2]int{i, j}] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix); k++ {
					for l := 0; l < len(matrix[k]); l++ {
						if (k == i || l == j) && matrix[k][l] != 0 {
							delete(m, [2]int{k, l})
							matrix[k][l] = math.MinInt32
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MinInt32 && m[[2]int{i, j}] == false {
				matrix[i][j] = 0
			}
		}
	}
}
