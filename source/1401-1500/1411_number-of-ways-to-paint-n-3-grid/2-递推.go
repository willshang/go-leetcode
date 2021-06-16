package main

import "fmt"

func main() {
	fmt.Println(numOfWays(10))
}

// leetcode1411_给Nx3网格图涂色的方案数
var mod = 1000000007

func numOfWays(n int) int {
	// 满足要求的组合12种
	// 6种互不相同：012, 021, 102, 120, 201, 210
	// 6种带有相同：010, 020, 101, 121, 202, 212
	a := 6 // 第1行是互不相同的
	b := 6 // 第1行带相同的
	for i := 2; i <= n; i++ {
		x := (2*a + 2*b) % mod // 当前行是互不相同的
		y := (2*a + 3*b) % mod // 当前行是带相同的
		a = x
		b = y
	}
	return (a + b) % mod
}
