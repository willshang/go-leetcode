package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPerfectNumber(24))
	fmt.Println(checkPerfectNumber(28))
}

// leetcode507_完美数
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i <= num/i; i++ {
		if num%i == 0 {
			sum = sum + i + (num / i)
		}
	}
	return sum == num
}
