package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countBits(10))
}

func countBits(num int) []int {
	res := make([]int, 0)
	for i := 0; i <= num; i++ {
		count := bits.OnesCount(uint(i))
		res = append(res, count)
	}
	return res
}
