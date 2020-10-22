package main

import "fmt"

func main() {
	fmt.Println(canPermutePalindrome("tactcoa"))
}

// 程序员面试金典01.04_回文排列
func canPermutePalindrome(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] == 2 {
			delete(m, s[i])
		}
	}
	return len(m) <= 1
}
