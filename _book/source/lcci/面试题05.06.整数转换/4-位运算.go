package main

import (
	"fmt"
)

func main() {
	//fmt.Println(convertInteger(1,2))
	fmt.Println(convertInteger(826966453, -729934991))
}

// 程序员面试金典05.06_整数转换
func convertInteger(A int, B int) int {
	C := A ^ B
	res := 0
	for i := 0; i < 32; i++ {
		if C&1 == 1 {
			res++
		}
		C = C >> 1
	}
	return res
}
