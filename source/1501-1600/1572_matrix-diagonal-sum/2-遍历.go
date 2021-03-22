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

func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		a, b := i, len(mat)-1-i
		res = res + mat[a][a]
		res = res + mat[a][b]

	}
	if len(mat)%2 == 1 {
		res = res - mat[len(mat)/2][len(mat)/2]
	}
	return res
}
