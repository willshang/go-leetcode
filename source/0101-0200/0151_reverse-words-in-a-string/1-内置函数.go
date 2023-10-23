package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

// leetcode151_翻转字符串里的单词
func reverseWords(s string) string {
	arr := strings.Fields(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, " ")
}
