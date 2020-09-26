package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(arr, 4, 1))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])
	if (row*col != r*c) || (row == r && col == c) {
		return nums
	}
	res := make([][]int, 0)
	arr := make([]int, 0)
	count := 0
	for _, num := range nums {
		for _, value := range num {
			arr = append(arr, value)
			count++
			if count == c {
				res = append(res, arr)
				arr = []int{}
				count = 0
			}
		}
	}
	return res
}
