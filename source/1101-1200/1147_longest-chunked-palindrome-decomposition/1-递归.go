package main

import "fmt"

func main() {
	fmt.Println(longestDecomposition("elvtoelvto"))
}

// leetcode1147_段式回文
func longestDecomposition(text string) int {
	n := len(text)
	if n <= 1 {
		return n
	}
	for i := 1; i <= n/2; i++ {
		if text[:i] == text[n-i:] {
			return 2 + longestDecomposition(text[i:n-i])
		}
	}
	return 1
}
