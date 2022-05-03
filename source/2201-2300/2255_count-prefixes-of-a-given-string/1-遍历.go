package main

import "strings"

func main() {

}

// leetcode2255_统计是给定字符串前缀的字符串数目
func countPrefixes(words []string, s string) int {
	res := 0
	for _, v := range words {
		if strings.HasPrefix(s, v) == true {
			res++
		}
	}
	return res
}
