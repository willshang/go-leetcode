package main

import "fmt"

func main() {
	fmt.Println(fib(100))
}

// 剑指offer_面试题10- I.斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, (n1+n2)%1000000007
	}
	return n2
}
