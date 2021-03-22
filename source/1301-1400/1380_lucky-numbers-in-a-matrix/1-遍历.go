package main

import "fmt"

func main() {
	arr := [][]int{
		{7, 8},
		{1, 2},
	}
	fmt.Println(luckyNumbers(arr))
}

// leetcode1380_矩阵中的幸运数
func luckyNumbers(matrix [][]int) []int {
	res := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		minIndex := 0
		for j := 1; j < len(matrix[i]); j++ {
			if min > matrix[i][j] {
				min = matrix[i][j]
				minIndex = j
			}
		}
		flag := true
		for j := 0; j < len(matrix); j++ {
			if matrix[j][minIndex] > min {
				flag = false
				break
			}
		}
		if flag == true {
			res = append(res, min)
		}
	}
	return res
}
