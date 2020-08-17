package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	str := strconv.FormatInt(int64(num), 4)
	return str[0:1] == "1" && strings.Count(str, "0") == len(str)-1
}
