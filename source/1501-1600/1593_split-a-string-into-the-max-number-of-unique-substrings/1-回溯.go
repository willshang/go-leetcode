package main

import "fmt"

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
}

// leetcode1593_拆分字符串使唯一子字符串的数目最大
var res int

func maxUniqueSplit(s string) int {
	res = 1
	dfs(s, make(map[string]bool), make([]string, 0))
	return res
}

func dfs(s string, m map[string]bool, arr []string) {
	if len(s) == 0 {
		if len(arr) > res {
			res = len(arr)
		}
		return
	}
	for i := 0; i < len(s); i++ {
		newStr := s[:i+1]
		if m[newStr] == false {
			m[newStr] = true
			dfs(s[i+1:], m, append(arr, newStr))
			m[newStr] = false
		}
	}
}
