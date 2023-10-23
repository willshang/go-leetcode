package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
}

// leetcode856_括号的分数
func scoreOfParentheses(S string) int {
	res := 0
	count := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			count++
		} else {
			count--
			if S[i-1] == '(' {
				res = res + 1<<count
			}
		}
	}
	return res
}
