package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(arr, 4, 1))
}

// leetcode566_重塑矩阵
func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])
	if (row*col != r*c) || (row == r && col == c) {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = nums[i/col][i%col]
	}
	return res
}
