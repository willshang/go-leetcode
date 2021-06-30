package main

import "strings"

func main() {

}

// leetcode1910_删除一个字符串中所有出现的给定子字符串
func removeOccurrences(s string, part string) string {
	for strings.Contains(s, part) {
		s = strings.Replace(s, part, "", 1)
	}
	return s
}
