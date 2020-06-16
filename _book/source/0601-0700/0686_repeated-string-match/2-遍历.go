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
	temp := A
	count := 1
	for len(temp) < len(B) {
		temp = temp + A
		count++
	}
	if strings.Contains(temp, B) {
		return count
	}
	temp = temp + A
	if strings.Contains(temp, B) {
		return count + 1
	}
	return -1
}
