package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	count := 1
	level := 1
	for count <= n*n {
		top, bottom, left, right := level-1, n-level, level-1, n-level
		// 左到右
		for i := left; i <= right && left <= right; i++ {
			res[top][i] = count
			count++
		}
		// 上到下
		for i := top + 1; i <= bottom && top <= bottom; i++ {
			res[i][right] = count
			count++
		}
		// 右到左
		for i := right - 1; i >= left && left <= right; i-- {
			res[bottom][i] = count
			count++
		}
		// 下到上
		for i := bottom - 1; i >= top+1 && top <= bottom; i-- {
			res[i][left] = count
			count++
		}
		level++
	}
	return res
}
