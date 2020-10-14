package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestSubstring("ababacb", 3))
}

func longestSubstring(s string, k int) int {
	res := 0
	m := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		m[rune(s[i])]++
	}
	split := make([]rune, 0)
	for key, value := range m {
		if value < k {
			split = append(split, key)
		}
	}
	if len(split) == 0 {
		return len(s)
	}
	arr := strings.FieldsFunc(s, func(r rune) bool {
		return m[r] < k
	})
	fmt.Println(arr, m)
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > res {
			res = len(arr[i])
		}
	}
	return res
}
