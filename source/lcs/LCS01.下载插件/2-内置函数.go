package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(leastMinutes(8))
}

func leastMinutes(n int) int {
	return bits.Len(uint(n)-1) + 1
}
