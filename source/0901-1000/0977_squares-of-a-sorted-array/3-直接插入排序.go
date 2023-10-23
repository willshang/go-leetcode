package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	res[0] = A[0] * A[0]
	j := 0
	for i := 1; i < len(A); i++ {
		value := A[i] * A[i]
		for j = i - 1; j >= 0; j-- {
			if value < res[j] {
				res[j+1] = res[j]
			} else {
				break
			}
		}
		res[j+1] = value
	}
	return res
}
