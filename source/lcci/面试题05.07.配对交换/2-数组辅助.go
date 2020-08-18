package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(exchangeBits(158))
}

func exchangeBits(num int) int {
	a := fmt.Sprintf("%b", num)
	arr := make([]byte, 32)
	for i := 0; i < 32; i++ {
		arr[i] = '0'
	}
	count := 31
	for i := len(a) - 1; i >= 0; i-- {
		arr[count] = a[i]
		count--
	}
	for i := len(arr) - 2; i >= 0; i = i - 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	value, _ := strconv.ParseInt(string(arr), 2, 64)
	return int(value)
}
