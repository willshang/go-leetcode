package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(53))
}


func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n % 3 == 0{
		n = n / 3
	}
	return n == 1
}