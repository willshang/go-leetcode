package main

import "fmt"

func main() {
	str := "A man, a plan, a canal: Panama"
	reverseString([]byte(str))
}

// leetcode344_反转字符串
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
