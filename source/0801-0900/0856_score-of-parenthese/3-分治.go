package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
}

func scoreOfParentheses(S string) int {
	return dfs(S, 0, len(S))
}

func dfs(S string, left, right int) int {
	res := 0
	count := 0
	for i := left; i < right; i++ {
		if S[i] == '(' {
			count++
		} else {
			count--
		}
		if count == 0 {
			if i-left == 1 {
				res++
			} else {
				res = res + 2*dfs(S, left+1, i)
			}
			left = i + 1
		}
	}
	return res
}
