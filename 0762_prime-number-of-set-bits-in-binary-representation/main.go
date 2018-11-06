package main

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(10,15))
}
func countPrimeSetBits(L int, R int) int {
	primes := [...]int{2:1,3:1,5:1,7:1,11:1,13:1,17:1,19:1}

	res := 0
	for i := L; i <= R; i++{
		bits := 0
		for n := i; n > 0; n >>= 1{
			bits = bits + n & 1
		}
		res = res + primes[bits]
	}
	return res
}