package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}

func countPrimeSetBits(L int, R int) int {
	primes := map[int]int{
		2:  1,
		3:  1,
		5:  1,
		7:  1,
		11: 1,
		13: 1,
		17: 1,
		19: 1,
		23: 1,
		29: 1,
		31: 1,
	}
	res := 0
	for i := L; i <= R; i++ {
		count := bits.OnesCount32(uint32(i))
		res = res + primes[count]
	}
	return res
}
