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
	for start, end := 0, n-1; start < end; {
		for s, e := start, end; s < end; {
			matrix[start][s], matrix[e][start], matrix[end][e], matrix[s][end] =
				matrix[e][start], matrix[end][e], matrix[s][end], matrix[start][s]
			s++
			e--
		}
		start++
		end--
	}
}
