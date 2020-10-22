package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	fmt.Println(a)
}

func merge(A []int, m int, B []int, n int) {
	A = A[:m]
	A = append(A, B[:n]...)
	sort.Ints(A)
}
