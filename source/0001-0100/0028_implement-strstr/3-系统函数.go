package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("hello", ""))
}

// leetcode 28 实现strStr()
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
