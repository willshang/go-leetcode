package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 2}
	fmt.Println(isMonotonic(arr))
}

func isMonotonic(A []int) bool {
	return inc(A) || desc(A)
}

func inc(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func desc(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			return false
		}
	}
	return true
}
