package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("aa"))
}

func maxLengthBetweenEqualCharacters(s string) int {
	res := -1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				res = max(res, j-i-1)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
