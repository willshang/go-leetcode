package main

import (
	"fmt"
)

func main() {
	//fmt.Println(convertInteger(1,2))
	fmt.Println(convertInteger(826966453, -729934991))
}

func convertInteger(A int, B int) int {
	C := uint32(A) ^ uint32(B)
	res := 0
	for C != 0 {
		res++
		C = C & (C - 1)
	}
	return res
}
