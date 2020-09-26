package main

import (
	"fmt"
	"math/bits"
)

func main() {
	//fmt.Println(convertInteger(1,2))
	fmt.Println(convertInteger(826966453, -729934991))
}

func convertInteger(A int, B int) int {
	C := uint32(A) ^ uint32(B)
	return bits.OnesCount(uint(C))
}
