package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	count := 0
	// 2 3 5 7
	for i := 5; i <= n; i = i + 2 {
		if isPrime(i) && isPrime(i-2) {
			count++
		}
	}
	fmt.Println(count)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
