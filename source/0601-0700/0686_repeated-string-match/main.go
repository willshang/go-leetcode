package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "abcd"
	B := "cdabcdab"
	fmt.Println(repeatedStringMatch(A, B))
	fmt.Println(repeatedStringMatch("abc", "cabcabca"))
}

func repeatedStringMatch(A string, B string) int {
	times := max(len(B)/len(A), 1)
	for i := 1; i <= times+2; i++ {
		if strings.Contains(strings.Repeat(A, i), B) {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
