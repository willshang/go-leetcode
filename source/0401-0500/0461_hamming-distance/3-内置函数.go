package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(hammingDistance(4, 1))
}

func hammingDistance(x int, y int) int {
	x = x ^ y
	return bits.OnesCount(uint(x))
}
