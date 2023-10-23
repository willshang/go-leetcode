package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(consecutiveNumbersSum(5))
}

func consecutiveNumbersSum(N int) int {
	res := 1
	// N=(x+1)+(x+2)+⋯+(x+k) = kx+k*(k+1)/2
	// 2N=k(2x+k+1)
	target := int(math.Sqrt(float64(2 * N)))
	for i := 1; i < target; i++ {
		left := N - i*(i+1)/2
		if left%(i+1) == 0 { // 长度i+1
			res++
		}
	}
	return res
}
