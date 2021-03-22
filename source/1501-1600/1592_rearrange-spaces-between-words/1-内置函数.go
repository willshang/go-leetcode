package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
}

// leetcode1592_重新排列单词间的空格
func reorderSpaces(text string) string {
	count := strings.Count(text, " ")
	arr := strings.Fields(text)
	if len(arr) == 1 {
		return arr[0] + strings.Repeat(" ", count)
	}
	value := count / (len(arr) - 1)
	left := count % (len(arr) - 1)
	res := strings.Join(arr, strings.Repeat(" ", value))
	res = res + strings.Repeat(" ", left)
	return res
}
