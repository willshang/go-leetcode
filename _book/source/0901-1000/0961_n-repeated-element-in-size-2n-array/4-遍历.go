package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}

// leetcode961_重复N次的元素
func repeatedNTimes(A []int) int {
	for i := 0; i < len(A)-2; i++ {
		if A[i] == A[i+1] || A[i] == A[i+2] {
			return A[i]
		}
	}
	return A[len(A)-1]
}
