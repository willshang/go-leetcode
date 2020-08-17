package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	for num >= 10 {
		num = sumDigits(num)
	}
	return num
}

func sumDigits(num int) int {
	sumVal := 0
	str := strconv.Itoa(num)
	for i := range str {
		sumVal = sumVal + int(str[i]-'0')
	}
	return sumVal
}
