package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1
	for i < j {
		if A[i]%2 == 0 {
			i++
		} else if A[j]%2 == 1 {
			j--
		} else {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
