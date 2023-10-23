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
	for a := 1; a <= n; a++ { // 上边界
		for b := a; b <= n; b++ { // 下边界
			left := 1
			value := 0
			for right := 1; right <= m; right++ {
				value = arr[b][right] - arr[b][left-1] - arr[a-1][right] + arr[a-1][left-1]
				if value > maxValue {
					maxValue = value
					res = []int{a - 1, left - 1, b - 1, right - 1}
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
