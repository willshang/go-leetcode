package main

import (
	"fmt"
)

func main() {
	fmt.Println(countDigitOne(12))
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	res := 0
	for i := 1; i <= n; i = i * 10 {
		left := n / i
		right := n % i
		res = res + (left+8)/10*i
		if left%10 == 1 {
			res = res + right + 1
		}
	}
	return res
}
