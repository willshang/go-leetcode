package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(diagonalSum(arr))
}

// leetcode1572_矩阵对角线元素的和
func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		a, b := i, len(mat)-1-i
		if a == b {
			res = res + mat[a][b]
		} else {
			res = res + mat[a][a]
			res = res + mat[a][b]
		}
	}
	return res
}
