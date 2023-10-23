package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(insertBits(0, 31, 0, 4))
	//fmt.Println(insertBits(1024, 19, 2, 6))
	//fmt.Println(insertBits(1143207437, 1017033, 11, 31))
	fmt.Println(insertBits(2032243561, 10, 24, 29))
}

func insertBits(N int, M int, i int, j int) int {
	arr := make([]byte, 32)
	for i := 0; i < 32; i++ {
		arr[i] = '0'
	}
	a := fmt.Sprintf("%b", N)
	b := fmt.Sprintf("%b", M)
	for k := len(a) - 1; k >= 0; k-- {
		arr[31-(len(a)-1-k)] = a[k]
	}
	count := 0
	for k := 31 - i; k >= 31-j; k-- {
		if count < len(b) {
			arr[k] = b[len(b)-1-count]
			count++
		} else {
			arr[k] = '0'
		}
	}
	value, _ := strconv.ParseInt(string(arr), 2, 64)
	return int(value)
}
