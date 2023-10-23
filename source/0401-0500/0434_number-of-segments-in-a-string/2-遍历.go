package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
}

// leetcode434_字符串中的单词数
func countSegments(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			count++
		}
	}
	return count
}
