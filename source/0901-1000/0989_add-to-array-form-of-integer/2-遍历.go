package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}

func addToArrayForm(A []int, K int) []int {
	A[len(A)-1] = A[len(A)-1] + K
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		carry = A[i] / 10
		A[i] = A[i] % 10
		if i > 0 {
			A[i-1] = A[i-1] + carry
		}
	}
	for carry > 0 {
		A = append([]int{carry % 10}, A...)
		carry = carry / 10
	}
	return A
}
