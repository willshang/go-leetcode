package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(thousandSeparator(987))
	fmt.Println(thousandSeparator(1234))
	fmt.Println(thousandSeparator(123456789))
	fmt.Println(thousandSeparator(0))
}

// leetcode1556_千位分隔数
func thousandSeparator(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	count := 0
	for n != 0 {
		count++
		value := n % 10
		if count%3 == 1 {
			res = strconv.Itoa(value) + "." + res
		} else {
			res = strconv.Itoa(value) + res
		}
		n = n / 10
	}
	return strings.Trim(res, ".")
}
