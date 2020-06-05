package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(4))
}

// n (10|01)&3(11)=10|01
func hasAlternatingBits(n int) bool {
	temp := n & 3
	if temp != 1 && temp != 2 {
		return false
	}
	for n > 0 {
		if n&3 != temp {
			return false
		}
		n = n >> 2
	}
	return true
}
