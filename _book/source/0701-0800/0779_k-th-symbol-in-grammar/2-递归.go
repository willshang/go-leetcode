package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(kthGrammar(4, 5))
}

// leetcode779_第K个语法符号
func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	total := int(math.Pow(2, float64(N-1)))
	half := total / 2
	if K <= half {
		return kthGrammar(N-1, K)
	}
	return 1 - kthGrammar(N-1, K-half)
}
