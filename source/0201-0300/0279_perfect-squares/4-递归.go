package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
}

var m map[int]int

func numSquares(n int) int {
	m = make(map[int]int)
	return dfs(n)
}

func dfs(n int) int {
	if m[n] > 0 {
		return m[n]
	}
	if n == 0 {
		return 0
	}
	count := math.MaxInt32
	for i := 1; i*i <= n; i++ {
		count = min(count, dfs(n-i*i)+1)
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
