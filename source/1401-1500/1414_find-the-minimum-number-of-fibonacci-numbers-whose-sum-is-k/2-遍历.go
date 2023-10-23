package main

import "fmt"

func main() {
	fmt.Println(findMinFibonacciNumbers(19))
}

func findMinFibonacciNumbers(k int) int {
	res := 0
	for {
		target := get(k)
		k = k - target
		res++
		if k == 0 {
			break
		}
	}
	return res
}

func get(k int) int {
	a, b := 1, 1
	for {
		a, b = b, a+b
		if b > k {
			return a
		} else if b == k {
			return b
		}
	}
}
