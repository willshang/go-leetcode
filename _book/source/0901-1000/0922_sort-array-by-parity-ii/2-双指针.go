package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	for i < len(A) {
		for A[i]%2 != 0 {
			if A[j]%2 == 0 {
				A[i], A[j] = A[j], A[i]
			} else {
				j = j + 2
			}
		}
		i = i + 2
	}
	return A
}
