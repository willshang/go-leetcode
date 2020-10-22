package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(4))
}

func hasAlternatingBits(n int) bool {
	str := strconv.FormatInt(int64(n), 2)
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return false
		}
	}
	return true
}
