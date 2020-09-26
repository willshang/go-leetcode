package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 2}
	fmt.Println(isMonotonic(arr))
}

// leetcode896_单调数列
func isMonotonic(A []int) bool {
	toEnd := true
	toFirst := true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			toEnd = false
		}
		if A[i] < A[i+1] {
			toFirst = false
		}
	}
	return toEnd || toFirst
}
