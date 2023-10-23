package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(30))
}

// a = 1010
// 4 = 0100
// a & 4 = 0
func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	// return (num & (num-1) == 0) && (num-1)%3 == 0
	return (num&(num-1) == 0) && (num&0xaaaaaaaa == 0)
}
