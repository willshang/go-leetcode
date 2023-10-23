package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

// 剑指OfferII098.路径的数目
// 求C((m+n-2),(m-1))=> (n+m-2)!/(m-1)!(n-1)!
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if m > n {
		m, n = n, m
	}
	a := 1
	for i := 1; i <= m-1; i++ {
		a = a * i
	}
	b := 1
	for i := n; i <= m+n-2; i++ {
		b = b * i
	}
	return b / a
}
