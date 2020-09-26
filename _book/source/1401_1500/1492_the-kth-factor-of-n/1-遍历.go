package main

import "fmt"

func main() {
	fmt.Println(kthFactor(12, 3))
}

// leetcode1492_n的第k个因子
func kthFactor(n int, k int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
			if count == k {
				return i
			}
		}
	}
	return -1
}
