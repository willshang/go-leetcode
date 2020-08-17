package main

import (
	"fmt"
)

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 1, 7}))
}

func minIncrementForUnique(A []int) int {
	res := 0
	arr := make([]int, 80001)
	for i := 0; i < len(A); i++ {
		arr[A[i]]++
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 2 {
			sum = sum + arr[i] - 1
			res = res - i*(arr[i]-1)
		} else if arr[i] == 0 && sum > 0 {
			sum--
			res = res + i
		}
	}
	return res
}
