package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

// leetcode59_螺旋矩阵II
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	count := 1
	top, bottom, left, right := 0, n-1, 0, n-1
	for count <= n*n {
		for i := left; i <= right; i++ {
			res[top][i] = count
			count++
		}
		top++
		for i := top; i <= bottom; i++ {
			res[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- {
			res[bottom][i] = count
			count++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res[i][left] = count
			count++
		}
		left++
	}
	return res
}
