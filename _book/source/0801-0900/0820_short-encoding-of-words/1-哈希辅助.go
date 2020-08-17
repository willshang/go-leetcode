package main

import "fmt"

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}

// leetcode820_单词的压缩编码
func minimumLengthEncoding(words []string) int {
	res := 0
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		m[words[i]] = true
	}
	for k := range m {
		for i := 1; i < len(k); i++ {
			delete(m, k[i:])
		}
	}
	for k := range m {
		res = res + len(k) + 1
	}
	return res
}
