package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4}))
	fmt.Println(validMountainArray([]int{4, 3, 2, 1, 0}))
}

// leetcode941_有效的山脉数组
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	j := len(A) - 1
	for i < j && A[i] < A[i+1] {
		i++
	}
	for i < j && A[j] < A[j-1] {
		j--
	}
	return i == j && i != 0 && j != len(A)-1
}
