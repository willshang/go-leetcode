package main

import "fmt"

func main() {
	fmt.Println(countPrimes(3))
}

// 把N的倍数去掉(厄拉多塞筛法)
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	notPrimes := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}
		for j := i * 2; j < n; j += i {
			notPrimes[j] = true
		}
		count++
	}
	return count
}
