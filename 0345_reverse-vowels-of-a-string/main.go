package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("hello"))
}
func reverseVowels(s string) string {
	bytes := []byte(s)
	i, j := 0, len(s)-1
	for {
		for i < len(s) && !isvowels(bytes[i]){
			i++
		}

		for 0 <= j && !isvowels(bytes[j]){
			j--
		}

		if i >= j{
			break
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
func isvowels(b byte)bool  {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		   b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}