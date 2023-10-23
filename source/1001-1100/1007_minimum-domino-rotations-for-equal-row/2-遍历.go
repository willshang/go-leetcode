package main

import "fmt"

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
}

// leetcode1007_行相等的最少多米诺旋转
func minDominoRotations(A []int, B []int) int {
	var arr, arrA, arrB [7]int
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			arr[A[i]]++
		} else {
			arr[A[i]]++
			arr[B[i]]++
			arrA[A[i]]++
			arrB[B[i]]++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == len(A) {
			return min(arrA[i], arrB[i])
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
