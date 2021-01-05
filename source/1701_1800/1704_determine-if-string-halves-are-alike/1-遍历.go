package main

import "strings"

func main() {

}

// leetcode1704_判断字符串的两半是否相似
func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	total := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) == true {
			if i < len(s)/2 {
				total++
			} else {
				total--
			}
		}
	}
	return total == 0
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' ||
		b == 'i' || b == 'o' || b == 'u'
}
