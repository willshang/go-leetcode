package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

// leetcode313_超级丑数
func nthSuperUglyNumber(n int, primes []int) int {
	arr := make([]int, n)
	arr[0] = 1
	times := make([]int, len(primes))
	for i := 1; i < n; i++ {
		next := math.MaxInt32
		for j, value := range times {
			next = min(next, primes[j]*arr[value])
		}
		for j, value := range times {
			if primes[j]*arr[value] == next {
				times[j]++
			}
		}
		arr[i] = next
	}
	return arr[n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
