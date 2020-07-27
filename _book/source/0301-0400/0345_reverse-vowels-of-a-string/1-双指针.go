package main

import (
	"fmt"
)

func main() {
	//fmt.Println(reverseVowels("leetcode"))
	//fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels(".,"))
}

// leetcode345_反转字符串中的元音字母
func reverseVowels(s string) string {
	bytes := []byte(s)
	length := len(s)
	i, j := 0, length-1
	for i < j {
		if !isvowels(bytes[i]) {
			i++
			continue
		}
		if !isvowels(bytes[j]) {
			j--
			continue
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func isvowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
