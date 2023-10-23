package main

import "math"

func main() {

}

// leetcode1975_最大方阵和
func maxMatrixSum(matrix [][]int) int64 {
	res := int64(0)
	minValue := math.MaxInt32
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res = res + int64(abs(matrix[i][j]))
			minValue = min(minValue, abs(matrix[i][j]))
			if matrix[i][j] <= 0 {
				count++
			}
		}
	}
	if count%2 == 0 {
		return res
	}
	return res - 2*int64(minValue)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
