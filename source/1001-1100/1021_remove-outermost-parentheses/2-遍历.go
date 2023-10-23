package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses(S string) string {
	res := ""
	count := 0
	last := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			count++
		} else {
			count--
		}
		if count == 1 && S[i] == '(' {
			last = i
		}
		if count == 0 {
			res = res + S[last+1:i]
		}
	}
	return res
}
