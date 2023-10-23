package main

import (
	"fmt"
)

func main() {
	//fmt.Println(findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
	fmt.Println(findLongestWord("aaa", []string{"aaa", "aa", "a"}))
}

func findLongestWord(s string, dictionary []string) string {
	res := ""
	for i := 0; i < len(dictionary); i++ {
		if isSubsequence(dictionary[i], s) {
			if len(dictionary[i]) > len(res) ||
				(len(dictionary[i]) == len(res) && dictionary[i] < res) {
				res = dictionary[i]
			}
		}
	}
	return res
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
