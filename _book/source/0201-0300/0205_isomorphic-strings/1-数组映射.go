package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("foo", "baa"))
}

// leetcode205_同构字符串
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		a := int(s[i])
		b := int(t[i])
		if m1[a] != m2[b] {
			return false
		}
		m1[a] = i + 1
		m2[b] = i + 1
	}
	return true
}
