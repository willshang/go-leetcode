package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 18},
	}
	fmt.Println(luckyNumbers(arr))
}

func luckyNumbers(matrix [][]int) []int {
	res := make([]int, 0)
	minArr := make([]int, 0)
	maxArr := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		for j := 1; j < len(matrix[i]); j++ {
			if min > matrix[i][j] {
				min = matrix[i][j]
			}
		}
		minArr = append(minArr, min)
	}
	for i := 0; i < len(matrix[0]); i++ {
		max := matrix[0][i]
		for j := 1; j < len(matrix); j++ {
			if max < matrix[j][i] {
				max = matrix[j][i]
			}
		}
		maxArr = append(maxArr, max)
	}
	for i := 0; i < len(minArr); i++ {
		for j := 0; j < len(maxArr); j++ {
			if minArr[i] == maxArr[j] {
				res = append(res, minArr[i])
			}
		}
	}
	return res
}
