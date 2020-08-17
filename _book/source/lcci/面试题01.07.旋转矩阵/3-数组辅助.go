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

func rotate(matrix [][]int) {
	n := len(matrix)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[j][n-1-i] = matrix[i][j]
		}
	}
	copy(matrix, arr)
}
