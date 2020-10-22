package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

var m map[string]bool
var visited map[int]bool

func wordBreak(s string, wordDict []string) bool {
	m = make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	visited = make(map[int]bool)
	return wordbreak(s, 0)
}

func wordbreak(s string, start int) bool {
	if start == len(s) {
		return true
	}
	if _, ok := visited[start]; ok {
		return visited[start]
	}
	//递归
	for i := start; i < len(s); i++ {
		if _, ok := m[s[start:i+1]]; ok && wordbreak(s, i+1) {
			visited[start] = true
			return true
		}
	}
	visited[start] = false
	return false
}
