package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords(" the sky is blue "))
}

// 剑指offer_面试题58-I_翻转单词顺序
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	arr := strings.Fields(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, " ")
}
