package main

import "fmt"

func main() {
	fmt.Println(permutation("abc"))
}

var m map[string]bool

func permutation(s string) []string {
	m = make(map[string]bool)
	dfs(s, "")
	res := make([]string, 0)
	for key := range m {
		res = append(res, key)
	}
	return res
}

func dfs(s string, str string) {
	if len(s) == 0 {
		m[str] = true
	}
	for i := 0; i < len(s); i++ {
		temp := s[0:i] + s[i+1:]
		dfs(temp, str+string(s[i]))
	}
}
