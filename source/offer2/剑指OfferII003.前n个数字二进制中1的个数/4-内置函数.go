package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countBits(10))
}

func countBits(n int) []int {
	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		count := bits.OnesCount(uint(i))
		res = append(res, count)
	}
	return res
}
