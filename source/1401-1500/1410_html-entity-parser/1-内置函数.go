package main

import (
	"fmt"
	"html"
	"strings"
)

func main() {
	fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))
}

// leetcode1410_HTML实体解析器
func entityParser(text string) string {
	text = html.UnescapeString(text)
	return strings.ReplaceAll(text, "⁄", "/")
}
