package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
	fmt.Println(firstUniqChar("leetcode"))
}

func firstUniqChar(s string) int {
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
