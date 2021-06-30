package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
}

// leetcode524_通过删除字母匹配到字典里最长单词
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) > len(dictionary[j])
	})
	for i := 0; i < len(dictionary); i++ {
		if isSubsequence(dictionary[i], s) {
			return dictionary[i]
		}
	}
	return ""
}

// leetcode392.判断子序列
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
