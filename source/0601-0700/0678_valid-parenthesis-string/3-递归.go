package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkValidString("()"))
}

func checkValidString(s string) bool {
	return dfs(s, 0, 0)
}

func dfs(s string, index, count int) bool {
	if count < 0 {
		return false
	}
	for i := index; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			if count == 0 {
				return false
			}
			count--
		} else if s[i] == '*' {
			return dfs(s, i+1, count+1) || dfs(s, i+1, count-1) || dfs(s, i+1, count)
		}
	}
	return count == 0
}
