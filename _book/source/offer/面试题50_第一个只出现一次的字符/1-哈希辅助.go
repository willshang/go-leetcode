package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
	fmt.Println(firstUniqChar(""))
}

// 剑指offer_面试题50_第一个只出现一次的字符
func firstUniqChar(s string) byte {
	res := byte(' ')
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return res
}
