package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(4))
}

func hasAlternatingBits(n int) bool {
	pre := n & 1
	n = n >> 1
	for n > 0 {
		if n&1 == pre {
			return false
		}
		pre = n & 1
		n = n >> 1
	}
	return true
}
