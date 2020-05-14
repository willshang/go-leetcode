package main

import (
	"fmt"
)

func main() {
	A := "abcde"
	B := "cdeab"
	fmt.Println(rotateString(A, B))
}

// leetcode796_旋转字符串
func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A); i++ {
		A = A[1:] + string(A[0])
		if A == B {
			return true
		}
	}
	return false
}
