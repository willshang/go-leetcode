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
	if n == 1 {
		return true
	}
	if n%2 != 0 {
		return false
	}
	return isPowerOfTwo(n / 2)
}
