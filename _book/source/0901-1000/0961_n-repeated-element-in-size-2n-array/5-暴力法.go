package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}

func repeatedNTimes(A []int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] {
				return A[i]
			}
		}
	}
	return A[len(A)-1]
}
