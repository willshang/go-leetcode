package main

import "strings"

func main() {

}

// leetcode1967_作为子字符串出现在单词中的字符串数目
func numOfStrings(patterns []string, word string) int {
	res := 0
	for i := 0; i < len(patterns); i++ {
		if strings.Contains(word, patterns[i]) == true {
			res++
		}
	}
	return res
}
