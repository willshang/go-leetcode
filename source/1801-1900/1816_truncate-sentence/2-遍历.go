package main

import (
	"fmt"
)

func main() {
	s := "a b c d"
	fmt.Println(truncateSentence(s, 2))
}

// leetcode1816_截断句子
func truncateSentence(s string, k int) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
			if count == k {
				return s[:i]
			}
		}
	}
	return s
}
