package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}

func repeatedNTimes(A []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if _, ok := m[A[i]]; ok {
			return A[i]
		}
		m[A[i]]++
	}
	return 0
}
