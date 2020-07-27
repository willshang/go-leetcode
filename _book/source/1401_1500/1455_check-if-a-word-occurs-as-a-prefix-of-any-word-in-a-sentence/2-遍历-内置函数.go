package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
}

func isPrefixOfWord(sentence string, searchWord string) int {
	arr := strings.Fields(sentence)
	for k, v := range arr {
		if len(v) >= len(searchWord) {
			if v[:len(searchWord)] == searchWord {
				return k + 1
			}
		}
	}
	return -1
}
