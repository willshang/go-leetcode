package main

import (
	"fmt"
)

func main() {
	fmt.Println(integerReplacement(8))
}

var m map[int]int

func integerReplacement(n int) int {
	m = make(map[int]int)
	return dfs(n)
}

func dfs(n int) int {
	if m[n] > 0 {
		return m[n]
	}
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		m[n] = dfs(n/2) + 1
	} else {
		m[n] = min(dfs(n-1), dfs(n+1)) + 1
	}
	return m[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
