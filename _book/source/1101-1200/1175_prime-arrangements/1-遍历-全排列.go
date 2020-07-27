package main

import (
	"fmt"
)

func main() {
	fmt.Println(numPrimeArrangements(5))
}

// leetcode1175_质数排列
func numPrimeArrangements(n int) int {
	primeNum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeNum++
		}
	}
	a := 1
	for i := 2; i <= primeNum; i++ {
		a = a * i % 1000000007
	}
	for i := 2; i <= n-primeNum; i++ {
		a = a * i % 1000000007
	}
	return a
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
