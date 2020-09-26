package main

import (
	"fmt"
)

func main() {
	//fmt.Println(reverseVowels("leetcode"))
	//fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels(".,"))
}

func reverseVowels(s string) string {
	bytes := []byte(s)
	length := len(s)
	temp := make([]byte, 0)
	for i := 0; i < length; i++ {
		if isvowels(bytes[i]) {
			temp = append(temp, bytes[i])
		}
	}
	count := 0
	for i := 0; i < length; i++ {
		if isvowels(bytes[i]) {
			bytes[i] = temp[len(temp)-1-count]
			count++
		}
	}
	return string(bytes)
}

func isvowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
