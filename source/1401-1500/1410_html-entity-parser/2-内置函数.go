package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))
}

func entityParser(text string) string {
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&apos;", "'")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&frasl;", "/")
	text = strings.ReplaceAll(text, "&amp;", "&")
	return text
}
