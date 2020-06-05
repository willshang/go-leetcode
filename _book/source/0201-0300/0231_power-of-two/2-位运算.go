package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(218))
	fmt.Println(isPowerOfTwo(1024))
}

// leetcode231_2的幂
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0
}
