package main

import "fmt"

func main() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
}

// leetcode931_下降路径最小和
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			minValue := matrix[i-1][j]
			if j > 0 {
				minValue = min(minValue, matrix[i-1][j-1])
			}
			if j < n-1 {
				minValue = min(minValue, matrix[i-1][j+1])
			}
			matrix[i][j] = matrix[i][j] + minValue
		}
	}
	res := matrix[n-1][0]
	for i := 0; i < n; i++ {
		res = min(res, matrix[n-1][i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
