package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

// leetcode140_单词拆分II
var m map[string]bool
var visited map[int][]string

func wordBreak(s string, wordDict []string) []string {
	m = make(map[string]bool)
	visited = make(map[int][]string)
	for _, str := range wordDict {
		m[str] = true
	}
	return dfs(s, 0)
}

func dfs(s string, level int) []string {
	if str, ok := visited[level]; ok {
		return str
	}
	res := make([]string, 0)
	for i := level + 1; i <= len(s); i++ {
		if m[s[level:i]] {
			if i != len(s) {
				arr := dfs(s, i)
				for _, str := range arr {
					res = append(res, s[level:i]+" "+str)
				}
			} else {
				res = append(res, s[level:i])
			}
		}
	}
	visited[level] = res
	return res
}
