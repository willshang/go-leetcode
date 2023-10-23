package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMaxMatrix([][]int{
		{9, -8, 1, 3, -2},
		{-3, 7, 6, -2, 4},
		{6, -4, -4, 8, -7},
		//{1, -3},
		//{-4, 9},
	}))
}

func getMaxMatrix(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	maxValue := math.MinInt32
	res := make([]int, 0)
	for a := 0; a < n; a++ { // 上边界
		arr := make([]int, m)
		for b := a; b < n; b++ { // 下边界
			left := 0
			value := 0
			for right := 0; right < m; right++ {
				arr[right] = arr[right] + matrix[b][right]
				value = value + arr[right]
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
