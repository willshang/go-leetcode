package main

import "sort"

func main() {

}

// 程序员面试金典17.15_最长单词
var m map[string]bool

func longestWord(words []string) string {
	m = make(map[string]bool)
	n := len(words)
	for i := 0; i < n; i++ {
		m[words[i]] = true
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) > len(words[j])
	})
	for i := 0; i < n; i++ { // 从最长最小字典序的开始找
		m[words[i]] = false
		if dfs(words[i]) == true {
			return words[i]
		}
	}
	return ""
}

func dfs(str string) bool {
	if len(str) == 0 || m[str] == true {
		return true
	}
	for i := 1; i <= len(str); i++ {
		subStr := str[:i]
		if m[subStr] == true {
			if dfs(str[i:]) == true {
				return true
			}
		}
	}
	return false
}
