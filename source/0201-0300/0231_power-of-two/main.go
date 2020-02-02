package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(218))
	fmt.Println(isPowerOfTwo(1024))
}
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n = n / 2
	}
	return true
}
