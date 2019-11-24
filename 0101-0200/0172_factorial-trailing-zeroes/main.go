package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(5))
}

func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		n = n / 5
		res = res + n
	}
	return res
}
