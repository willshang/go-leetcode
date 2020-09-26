package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(53))
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	str := strconv.FormatInt(int64(n), 3)
	return str[0:1] == "1" && strings.Count(str, "0") == len(str)-1
}
