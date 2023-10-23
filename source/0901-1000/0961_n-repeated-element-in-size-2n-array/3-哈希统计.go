package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}

func repeatedNTimes(A []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m[A[i]]++
	}
	for key, value := range m {
		if value == len(A)/2 {
			return key
		}
	}
	return 0
}
