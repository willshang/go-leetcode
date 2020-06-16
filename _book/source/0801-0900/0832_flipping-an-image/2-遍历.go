package main

import "fmt"

func main() {
	arr := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(arr))
}

// leetcode832_翻转图像
func flipAndInvertImage(A [][]int) [][]int {
	for k := 0; k < len(A); k++ {
		i, j := 0, len(A[k])-1
		for i < j {
			A[k][i], A[k][j] = A[k][j], A[k][i]
			i++
			j--
		}
		for i := 0; i < len(A[k]); i++ {
			A[k][i] = A[k][i] ^ 1
		}
	}
	return A
}
