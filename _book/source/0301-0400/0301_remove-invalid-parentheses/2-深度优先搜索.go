package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}

var m map[string]bool
var max int

func removeInvalidParentheses(s string) []string {
	m = make(map[string]bool)
	res := make([]string, 0)
	max = 0
	dfs(s, 0, 0, "")
	for k := range m {
		res = append(res, k)
	}
	return res
}

func dfs(s string, start, count int, temp string) {
	if count < 0 {
		return
	}
	if start == len(s) {
		if count == 0 {
			if len(temp) > max {
				max = len(temp)
				m = make(map[string]bool)
				m[temp] = true
			} else if max == len(temp) {
				m[temp] = true
			}
		}
		return
	}
	if s[start] == '(' {
		dfs(s, start+1, count+1, temp+"(")
	} else if s[start] == ')' {
		dfs(s, start+1, count-1, temp+")")
	} else {
		dfs(s, start+1, count, temp+string(s[start]))
	}
	if s[start] == '(' || s[start] == ')' {
		dfs(s, start+1, count, temp)
	}
}
