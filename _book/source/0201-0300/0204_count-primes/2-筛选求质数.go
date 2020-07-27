package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

// leetcode204_计数质数
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	for i := 2; i*i < n; i++ {
		if !isPrimes[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrimes[j] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			count++
		}
	}
	return count
}
