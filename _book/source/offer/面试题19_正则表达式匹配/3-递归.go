package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(p) == 0 {
		return false
	}
	match := false
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		match = true
	}
	if len(p) > 1 && p[1] == '*' {
		return (match && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	return match && isMatch(s[1:], p[1:])
}
