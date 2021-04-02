package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDifferentIntegers("0a0"))
}

// leetcode1805_字符串中不同整数的数目
func numDifferentIntegers(word string) int {
	m := make(map[int]bool)
	for i := 0; i < len(word); i++ {
		if '0' <= word[i] && word[i] <= '9' {
			value := int(word[i] - '0')
			for i+1 < len(word) && '0' <= word[i+1] && word[i+1] <= '9' {
				i++
				value = value*10 + int(word[i]-'0')
			}
			m[value] = true
		}
	}
	return len(m)
}
