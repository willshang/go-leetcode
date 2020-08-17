package main

import (
	"fmt"
)

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}))
	//fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
}

var arr []int

func minIncrementForUnique(A []int) int {
	res := 0
	arr = make([]int, 80001)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
	for i := 0; i < len(A); i++ {
		b := findNext(A[i])
		res = res + b - A[i]
	}
	return res
}

func findNext(a int) int {
	b := arr[a]
	if b == -1 {
		arr[a] = a
		return a
	}
	b = findNext(b + 1)
	arr[a] = b
	return b
}
