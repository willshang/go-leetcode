package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}

// leetcode301_删除无效的括号
func removeInvalidParentheses(s string) []string {
	res := make([]string, 0)
	cur := make(map[string]int)
	cur[s] = 1
	for {
		for k := range cur {
			if isValid(k) {
				res = append(res, k)
			}
		}
		if len(res) > 0 {
			return res
		}
		next := make(map[string]int)
		for k := range cur {
			for i := range k {
				if k[i] == '(' || k[i] == ')' {
					str := k[:i] + k[i+1:]
					next[str] = 1
				}
			}
		}
		cur = next
	}
}

func isValid(s string) bool {
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			left--
		}
		if left < 0 {
			return false
		}
	}
	return left == 0
}
