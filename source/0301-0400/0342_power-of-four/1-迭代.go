package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(30))
}

// leetcode342_4的幂
func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	for num > 1 {
		if num%4 != 0 {
			return false
		}
		num = num / 4
	}
	return num == 1
}
