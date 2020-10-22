package main

import "fmt"

func main() {
	fmt.Println(fib(100))
}

var m = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if m[n] > 0 {
		return m[n]
	}
	m[n] = (fib(n-1) + fib(n-2)) % 1000000007
	return m[n]
}
