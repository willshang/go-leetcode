package main

import "fmt"

func main() {
	fmt.Println(countPrimes(100000))
}

// 大于2的偶数都不是质数，所以我们只需要判断奇数是不是质数。
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	isPrimes := make([]bool, n)
	count := n / 2
	//奇数
	for i := 3; i*i < n; i += 2 {
		fmt.Println(i, count)
		if isPrimes[i] {
			continue
		}
		for j := i * i; j < n; j = j + 2*i {
			if !isPrimes[j] {
				count--
				isPrimes[j] = true
			}
		}

	}
	return count
}
