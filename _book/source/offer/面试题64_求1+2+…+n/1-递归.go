package main

import "fmt"

func main() {
	fmt.Println(sumNums(10))
}

// 剑指offer_面试题64_求1+2+…+n
var res int

func sumNums(n int) int {
	res = 0
	dfs(n)
	return res
}

func dfs(n int) bool {
	res = res + n
	return n > 0 && dfs(n-1)
}
