package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

// leetcode10_正则表达式匹配
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(p) == 0 {
		return false
	}
	match := false
	// 正常匹配条件=>相等，或者 p[0]等于.就不用管s[0]
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		match = true
	}
	// 匹配多个 就把 s 往后移1位，注意p不移动
	// 匹配0个 就把 p 往后移2位，相当于p的*当前作废
	if len(p) > 1 && p[1] == '*' {
		return (match && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	// 匹配当前成功，同时往后移
	return match && isMatch(s[1:], p[1:])
}
