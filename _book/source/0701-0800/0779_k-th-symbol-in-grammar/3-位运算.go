package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(kthGrammar(4, 5))
}

func kthGrammar(N int, K int) int {
	return bits.OnesCount(uint(K-1)) % 2
}
