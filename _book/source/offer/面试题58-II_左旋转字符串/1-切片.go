package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

func reverseLeftWords(s string, n int) string {
	n = n % len(s)
	return s[n:] + s[:n]
}
