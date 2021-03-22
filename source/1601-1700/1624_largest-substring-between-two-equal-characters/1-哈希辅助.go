package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("aa"))
}

// leetcode1624_两个相同字符之间的最长子字符串
func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[byte]int)
	res := -1
	for i := 0; i < len(s); i++ {
		if value, ok := m[s[i]]; ok {
			res = max(res, i-value-1)
		} else {
			m[s[i]] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
