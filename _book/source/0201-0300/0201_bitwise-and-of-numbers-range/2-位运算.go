package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}

func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n = n & (n - 1) // n抹去右边1位1
	}
	return n
}
