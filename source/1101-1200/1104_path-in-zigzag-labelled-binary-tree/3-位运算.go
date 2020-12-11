package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(7 ^ 3)
	fmt.Println(pathInZigZagTree(14))
}

func pathInZigZagTree(label int) []int {
	res := make([]int, 0)
	for label > 1 {
		res = append([]int{label}, res...)
		label = label / 2
		length := bits.Len32(uint32(label)) - 1
		label = label ^ ((1 << length) - 1) // 7^3=4 => 111^11=100
	}
	res = append([]int{1}, res...)
	return res
}
