package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(30))
}

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	if num%4 != 0 {
		return false
	}

	return isPowerOfFour(num / 4)
}
