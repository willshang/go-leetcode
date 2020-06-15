package main

import (
	"fmt"
)

func main() {
	//fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
}

func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	for i := 0; i < len(A)-1; i++ {
		for j := 0; j < len(A)-1-i; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
		if i >= 2 {
			index := len(A) - 1 - i
			if A[index]+A[index+1] > A[index+2] {
				return A[index] + A[index+1] + A[index+2]
			}
		}
	}
	if A[0]+A[1] > A[2] {
		return A[0] + A[1] + A[2]
	}
	return 0
}
