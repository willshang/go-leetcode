package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{9, 5, 3, 3}))
}

func repeatedNTimes(A []int) int {
	sort.Ints(A)
	if A[len(A)/2] == A[len(A)/2+1] {
		return A[len(A)/2]
	} else {
		return A[len(A)/2-1]
	}
}
