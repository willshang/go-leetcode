package main

import "fmt"

func main() {
	fmt.Println(minInsertions("(())"))
}

func minInsertions(s string) int {
	res := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if i+1 < len(s) && s[i+1] == ')' {
				i++
			} else {
				res++ // 少一个')'补一个')'到2个')'
			}
			if left > 0 {
				left--
			} else {
				res++ // 左边无'('补一个
			}
		}
	}
	res = res + left*2
	return res
}
