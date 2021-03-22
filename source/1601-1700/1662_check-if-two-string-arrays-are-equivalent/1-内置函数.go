package main

import "strings"

func main() {

}

// leetcode1662_检查两个字符串数组是否相等
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
