package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

// leetcode537_复数乘法
func complexNumberMultiply(a string, b string) string {
	a1, a2 := getValue(a)
	b1, b2 := getValue(b)
	return fmt.Sprintf("%d+%di", a1*b1-a2*b2, a1*b2+a2*b1)
}

func getValue(str string) (a, b int) {
	arr := strings.Split(str, "+")
	a, _ = strconv.Atoi(arr[0])
	b, _ = strconv.Atoi(arr[1][:len(arr[1])-1])
	return a, b
}
