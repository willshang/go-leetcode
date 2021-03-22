package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}

func stringMatching(words []string) []string {
	res := make([]string, 0)
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[i], words[j]) {
				if _, ok := m[words[j]]; !ok {
					res = append(res, words[j])
					m[words[j]] = true
				}
			} else if strings.Contains(words[j], words[i]) {
				if _, ok := m[words[i]]; !ok {
					res = append(res, words[i])
					m[words[i]] = true
				}
			}
		}
	}
	return res
}
