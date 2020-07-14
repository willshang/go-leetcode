package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// leetcode22_括号生成
var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	dfs(0, 0, n, "")
	return res
}

func dfs(left, right, max int, str string) {
	if left == right && left == max {
		res = append(res, str)
		return
	}
	if left < max {
		dfs(left+1, right, max, str+"(")
	}
	if right < left {
		dfs(left, right+1, max, str+")")
	}
}
