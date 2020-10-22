package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res = append(res, A[i])
		}
	}
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 1 {
			res = append(res, A[i])
		}
	}
	return res
}
