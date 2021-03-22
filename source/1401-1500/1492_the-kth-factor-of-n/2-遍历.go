package main

import "fmt"

func main() {
	fmt.Println(kthFactor(7, 2))
}

func kthFactor(n int, k int) int {
	count := 0
	i := 1
	for i = 1; i*i <= n; i++ {
		if n%i == 0 {
			count++
			if count == k {
				return i
			}
		}
	}
	i--
	if i*i == n {
		i--
	}
	for ; i > 0; i-- {
		if n%i == 0 {
			count++
			if count == k {
				return n / i
			}
		}
	}
	return -1
}
