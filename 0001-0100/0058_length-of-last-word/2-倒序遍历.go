package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a "))
}

// leetcode 58 最后一个单词的长度
func lengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	result := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result > 0 {
				return result
			}
			continue
		}
		result++
	}
	return result
}
