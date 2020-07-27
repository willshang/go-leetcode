package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4}))
	fmt.Println(validMountainArray([]int{4, 3, 2, 1, 0}))
}

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	pre := A[0]
	i := 0
	for i = 1; i < len(A); i++ {
		if pre < A[i] {
			pre = A[i]
		} else if pre == A[i] {
			return false
		} else {
			break
		}
	}
	if i >= len(A) || i == 1 {
		return false
	}
	for ; i < len(A); i++ {
		if pre > A[i] {
			pre = A[i]
		} else if pre == A[i] {
			return false
		} else {
			return false
		}
	}
	return true
}
