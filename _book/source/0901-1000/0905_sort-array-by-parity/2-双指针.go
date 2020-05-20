package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1
	for i < j {
		for i < j && A[i]%2 == 0 {
			i++
		}
		for i < j && A[j]%2 == 1 {
			j--
		}
		A[i], A[j] = A[j], A[i]
	}
	return A
}
