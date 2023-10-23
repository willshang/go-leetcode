package main

import (
	"fmt"
	"sort"
)

func main() {
	//str := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	//fmt.Println(longestWord(str))
	str_1 := []string{"worlds", "w", "wo", "wor", "worl", "world", "a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(str_1))
}

// leetcode720_词典中最长的单词
func longestWord(words []string) string {
	if len(words) == 0 {
		return ""
	} else if len(words) == 1 && len(words[0]) > 1 {
		return ""
	}
	sort.Strings(words)
	m := make(map[string]bool)
	res := words[0]
	for _, w := range words {
		n := len(w)
		if n == 1 {
			m[w] = true
		} else if m[w[:n-1]] {
			m[w] = true
			if len(res) < len(w) {
				res = w
			}
		}
	}
	return res
}
