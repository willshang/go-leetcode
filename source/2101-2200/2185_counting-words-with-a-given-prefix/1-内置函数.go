package main

import "strings"

func main() {

}

// leetcode2185_统计包含给定前缀的字符串
func prefixCount(words []string, pref string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], pref) {
			res++
		}
	}
	return res
}
