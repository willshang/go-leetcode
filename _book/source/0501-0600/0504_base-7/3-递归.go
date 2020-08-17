package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(10))
	fmt.Println(convertToBase7(-7))
}

// leetcode504_七进制数
func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(-1*num)
	}
	if num < 7 {
		return strconv.Itoa(num)
	}
	return convertToBase7(num/7) + strconv.Itoa(num%7)
}
