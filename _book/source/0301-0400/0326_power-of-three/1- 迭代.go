package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(53))
}

// leetcode326_3的幂
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	return n == 1
}
