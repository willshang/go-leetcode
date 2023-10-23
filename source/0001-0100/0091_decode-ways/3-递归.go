package main

import "fmt"

func main() {
	fmt.Println(numDecodings("226"))
}

var m map[string]int

func numDecodings(s string) int {
	m = make(map[string]int)
	return dfs(s)
}

func dfs(s string) int {
	if m[s] > 0 {
		return m[s]
	}
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	if (s[0]-'0')*10+s[1]-'0' > 26 {
		return dfs(s[1:])
	}
	m[s] = dfs(s[1:]) + dfs(s[2:])
	return m[s]
}
