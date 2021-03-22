package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateTheString(2))
}

// leetcode1374_生成每种字符都是奇数个的字符串
func generateTheString(n int) string {
	if n%2 == 0 {
		return strings.Repeat("a", n-1) + "b"
	}
	return strings.Repeat("a", n)
}
