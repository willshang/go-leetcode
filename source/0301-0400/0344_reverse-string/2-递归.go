package main

import "fmt"

func main() {
	str := "A man, a plan, a canal: Panama"
	reverseString([]byte(str))
}

func reverseString(s []byte) {
	var reverse func(int, int)
	reverse = func(left, right int) {
		if left < right {
			s[left], s[right] = s[right], s[left]
			reverse(left+1, right-1)
		}
	}
	reverse(0, len(s)-1)
}
