package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
	fmt.Println(firstUniqChar(""))
}

func firstUniqChar(s string) byte {
	res := byte(' ')
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return res
}
