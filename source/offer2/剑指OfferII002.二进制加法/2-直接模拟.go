package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

// 剑指OfferII002.二进制加法
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	result := ""
	flag := 0
	current := 0

	for i >= 0 || j >= 0 {
		intA, intB := 0, 0
		if i >= 0 {
			intA = int(a[i] - '0')
		}
		if j >= 0 {
			intB = int(b[j] - '0')
		}
		current = intA + intB + flag
		flag = 0
		if current >= 2 {
			flag = 1
			current = current - 2
		}
		cur := strconv.Itoa(current)
		result = cur + result
		i--
		j--
	}
	if flag == 1 {
		result = "1" + result
	}
	return result
}
