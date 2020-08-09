package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	fmt.Println(a)
}

func merge(A []int, m int, B []int, n int) {
	temp := make([]int, m)
	copy(temp, A)

	if n == 0 {
		return
	}
	first, second := 0, 0
	for i := 0; i < len(A); i++ {
		if second >= n {
			A[i] = temp[first]
			first++
			continue
		}
		if first >= m {
			A[i] = B[second]
			second++
			continue
		}
		if temp[first] < B[second] {
			A[i] = temp[first]
			first++
		} else {
			A[i] = B[second]
			second++
		}
	}
}
