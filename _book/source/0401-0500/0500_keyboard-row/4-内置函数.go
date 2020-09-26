package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(str))
}

// leetcode500_键盘行
func findWords(words []string) []string {
	res := make([]string, 0, len(words))
	q := "qwertyuiopQWERTYUIOP"
	a := "asdfghjklASDFGHJKL"
	z := "zxcvbnmZXCVBNM"
	for _, word := range words {
		qLen, aLen, zLen := 0, 0, 0
		for i := 0; i < len(word); i++ {
			if strings.Contains(q, string(word[i])) {
				qLen++
			}
			if strings.Contains(a, string(word[i])) {
				aLen++
			}
			if strings.Contains(z, string(word[i])) {
				zLen++
			}
		}
		if qLen == len(word) || aLen == len(word) || zLen == len(word) {
			res = append(res, word)
		}
	}
	return res
}
