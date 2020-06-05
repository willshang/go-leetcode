package main

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}

// leetcode762_二进制表示中质数个计算置位
func countPrimeSetBits(L int, R int) int {
	primes := [...]int{
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
		count := 0
		for n := i; n > 0; n >>= 1 {
			// bits = bits + n & 1
			count = count + n%2
		}
		res = res + primes[count]
	}
	return res
}
