package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
	fmt.Println(minimumLengthEncoding([]string{"time", "btime", "atime"}))
}

func minimumLengthEncoding(words []string) int {
	res := 0
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		m[words[i]] = true
	}
	words = make([]string, 0)
	for k := range m {
		words = append(words, k)
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := len(words) - 1; i >= 0; i-- {
		if m[words[i]] == false {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if strings.HasSuffix(words[i], words[j]) == true {
				m[words[j]] = false
			}
		}
	}
	for k := range m {
		if m[k] == true {
			res = res + len(k) + 1
		}
	}
	return res
}
