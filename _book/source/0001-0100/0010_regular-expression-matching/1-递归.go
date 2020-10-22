package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

func isMatch(s string, p string) bool {
	return dfs(s, p, 0, 0)
}

func dfs(s string, p string, i, j int) bool {
	if i >= len(s) && j >= len(p) {
		return true
	}
	if i <= len(s) && j >= len(p) {
		return false
	}
	if j+1 < len(p) && p[j+1] == '*' {
		if (i < len(s) && p[j] == s[i]) || (p[j] == '.' && i < len(s)) {
			return dfs(s, p, i+1, j+2) ||
				dfs(s, p, i+1, j) ||
				dfs(s, p, i, j+2)
		} else {
			return dfs(s, p, i, j+2)
		}
	}
	if (i < len(s) && s[i] == p[j]) || (p[j] == '.' && i < len(s)) {
		return dfs(s, p, i+1, j+1)
	}
	return false
}
