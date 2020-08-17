package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
	fmt.Println(firstUniqChar("leetcode"))
}

func firstUniqChar(s string) int {
	for i := 0; i < len(s); i++ {
		flag := true
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
