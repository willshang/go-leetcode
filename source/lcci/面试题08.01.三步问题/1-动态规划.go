package main

import "fmt"

func main() {
	fmt.Println(waysToStep(3))
}

func waysToStep(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	a, b, c := 1, 2, 4
	for i := 4; i <= n; i++ {
		a, b, c = b, c, (a+b+c)%1000000007
	}
	return c
}
