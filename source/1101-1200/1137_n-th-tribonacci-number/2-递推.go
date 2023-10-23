package main

import "fmt"

func main() {
	fmt.Println(tribonacci(5))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a := 0
	b := 1
	c := 1
	for i := 3; i <= n; i++ {
		c, b, a = a+b+c, c, b
	}
	return c
}
