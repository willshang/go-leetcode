package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 2}
	fmt.Println(isMonotonic(arr))
}

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	temp := A[len(A)-1] - A[0]
	for i := 0; i < len(A)-1; i++ {
		if temp > 0 && A[i] > A[i+1] {
			return false
		} else if temp < 0 && A[i] < A[i+1] {
			return false
		} else if temp == 0 && A[i] != A[i+1] {
			return false
		}
	}
	return true
}
