package main

import "fmt"

func main() {
	fmt.Println(restoreMatrix([]int{3, 8}, []int{4, 7}))
}

// leetcode1605_给定行和列的和求可行矩阵
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n := len(rowSum)
	m := len(colSum)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			value := min(rowSum[i], colSum[j])
			res[i][j] = value
			rowSum[i] = rowSum[i] - value
			colSum[j] = colSum[j] - value
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
