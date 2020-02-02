package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arrs := []byte(strconv.Itoa(x))
	Len := len(arrs)
	for i := 0; i < Len/2; i++ {
		if arrs[i] != arrs[Len-i-1] {
			return false
		}
	}
	return true
}
