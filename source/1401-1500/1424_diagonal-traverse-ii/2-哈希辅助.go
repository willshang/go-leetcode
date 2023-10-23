package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(arr))
}

// leetcode1424_对角线遍历II
func findDiagonalOrder(nums [][]int) []int {
	maxValue := 0
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			m[i+j] = append(m[i+j], nums[i][j])
			if i+j > maxValue {
				maxValue = i + j
			}
		}
	}
	res := make([]int, 0)
	for i := 0; i <= maxValue; i++ {
		for j := len(m[i]) - 1; j >= 0; j-- {
			res = append(res, m[i][j])
		}
	}
	return res
}
