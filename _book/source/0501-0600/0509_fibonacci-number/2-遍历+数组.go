package main

import "fmt"

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	res := make([]int, N+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= N; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[N]
}
