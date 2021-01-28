package main

import "fmt"

func main() {
	fmt.Println(lexicalOrder(200))
}

// leetcode386_字典序排数
var res []int

func lexicalOrder(n int) []int {
	res = make([]int, 0)
	for i := 1; i <= 9; i++ {
		dfs(n, i)
	}
	return res
}

func dfs(n, start int) {
	if start > n {
		return
	}
	// N叉树前序遍历
	res = append(res, start)
	for i := 0; i <= 9; i++ {
		dfs(n, start*10+i)
	}
}
