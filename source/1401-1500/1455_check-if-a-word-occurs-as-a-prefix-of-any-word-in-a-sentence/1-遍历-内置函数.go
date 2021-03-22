package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
}

// leetcode1455_检查单词是否为句中其他单词的前缀
func isPrefixOfWord(sentence string, searchWord string) int {
	arr := strings.Split(sentence, " ")
	for k, v := range arr {
		if strings.HasPrefix(v, searchWord) {
			return k + 1
		}
	}
	return -1
}
