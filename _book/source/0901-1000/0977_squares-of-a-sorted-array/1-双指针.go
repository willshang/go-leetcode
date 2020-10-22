package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	i := 0
	j := len(A) - 1
	index := len(A) - 1
	for i <= j {
		if A[i]*A[i] < A[j]*A[j] {
			res[index] = A[j] * A[j]
			j--
		} else {
			res[index] = A[i] * A[i]
			i++
		}
		index--
	}
	return res
}
