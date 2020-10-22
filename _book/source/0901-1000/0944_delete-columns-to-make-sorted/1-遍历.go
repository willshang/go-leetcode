package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
}

// leetcode944_删列造序
func minDeletionSize(A []string) int {
	res := 0
	for i := 0; i < len(A[0]); i++ {
		for j := 1; j < len(A); j++ {
			if A[j][i] < A[j-1][i] {
				res++
				break
			}
		}
	}
	return res
}
