package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMaxMatrix([][]int{
		//{9, -8, 1, 3, -2},
		//{-3, 7, 6, -2, 4},
		//{6, -4, -4, 8, -7},
		{1, -3},
		{-4, 9},
	}))
}

// 程序员面试金典17.24_最大子矩阵
func getMaxMatrix(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[i+1][j+1] = matrix[i][j] + arr[i+1][j] + arr[i][j+1] - arr[i][j]
		}
	}
	maxValue := math.MinInt32
	res := make([]int, 0)
	for a := 0; a < n; a++ { // 上边界
		for b := a; b < n; b++ { // 下边界
			left := 0
			value := 0
			for right := 0; right < m; right++ {
				value = arr[b+1][right+1] - arr[b+1][left] - arr[a][right+1] + arr[a][left]
				if value > maxValue {
					maxValue = value
					res = []int{a, left, b, right}
				}
				if value < 0 {
					value = 0
					left = right + 1
				}
			}
		}
	}
	return res
}
