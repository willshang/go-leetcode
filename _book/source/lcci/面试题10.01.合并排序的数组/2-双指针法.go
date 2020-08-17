package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	fmt.Println(a)
}

// 程序员面试金典10.01_合并排序的数组
func merge(A []int, m int, B []int, n int) {
	for m > 0 && n > 0 {
		if A[m-1] < B[n-1] {
			A[m+n-1] = B[n-1]
			n--
		} else {
			A[m+n-1] = A[m-1]
			m--
		}
	}
	if m == 0 && n > 0 {
		for n > 0 {
			A[n-1] = B[n-1]
			n--
		}
	}
}
